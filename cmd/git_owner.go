package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var ownersToShow int
var commitsToQuery int

var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Determine owners for the given files",
	Long: `Ownership can be calculated in a few ways, some of which this command uses.

  The command evaluates the entire commit history and looks at the top three committers of all time.
  But it also presents the top three committers of the last n (default 50). People who have touched the file
  a lot can act as owners, but  if the file has changed a lot over time, the more recent contributors may know
  more about its current state.

  It also looks at the volume of changes (adds and deletes) across all commits and the last n. Someone who
  has made a lot of changes to the file is probably better informed about the structure than someone who made a bug fix
  on a line.

  It does not include comments from PRs, another common way to assess ownership, because that information is not
  available in the git repo.
  `,
	Args: cobra.MinimumNArgs(1),
	Run:  findOwners,
}

func init() {
	ownerCmd.Flags().IntVarP(&commitsToQuery, "query-size", "c", 50, "The number of commits to query for recent history")
	ownerCmd.Flags().IntVarP(&ownersToShow, "owner-count", "n", 3, "The number of owners to show in each list.")
	gitCmd.AddCommand(ownerCmd)
}

// ownerTotals contains the summation of touches and changes
// for a single user against a single file
type ownerTotals struct {
	owner        string
	totalTouches int
	totalChanges int // adds + deletes
}

func (totals *ownerTotals) String() string {
	return fmt.Sprintf("%s:%d;%d", totals.owner, totals.totalTouches, totals.totalChanges)
}

// updateTalliesForFile looks through the commit for references to fileName and updates
// the internal tallies appropriately.
func (totals *ownerTotals) updateTalliesForFile(fileName string, commit *object.Commit) error {
	changes, err := findChangesForFileInCommit(fileName, commit)
	if err != nil {
		return err
	}
	// this can happen if the commit is included even though it doesn't include the file
	// (a bug in the library maybe?)
	if changes != 0 {
		totals.totalTouches += 1
		totals.totalChanges += changes
	}
	return nil
}

// findChangesForFileInCommit uses commit to find the total number of changes
// for the given file in that commit.
func findChangesForFileInCommit(fileName string, commit *object.Commit) (int, error) {
	stats, err := commit.Stats()
	if err != nil {
		return 0, err
	}

	for _, fileStat := range stats {
		if fileStat.Name == fileName {
			return (fileStat.Addition + fileStat.Deletion), nil
		}
	}
	// while in theory this shouldn't happen because we specify the file name
	// we want from repo.Log, in practice commits can show up even if they don't
	// include the requested file (perhaps a bug in the library)
	return 0, nil
}

// a structure for keeping track of all the information about
// one dive into the commits
type fileLogQuery struct {
	tallies      map[string]*ownerTotals
	maxCommits   int
	description  string
	totalTouches int
	totalChanges int
}

// ownerTotals returns just a slice of the owner totals, to use in sorting
func (query *fileLogQuery) ownerTotals() []*ownerTotals {
	extracted := make([]*ownerTotals, 0, len(query.tallies))
	for _, tally := range query.tallies {
		extracted = append(extracted, tally)
	}
	return extracted
}

func (query *fileLogQuery) ensureOwnerTotals(owner string) *ownerTotals {
	authorTallies, exists := query.tallies[owner]
	if !exists {
		authorTallies = &ownerTotals{owner, 0, 0}
		query.tallies[owner] = authorTallies
		return authorTallies
	} else {
		return authorTallies
	}
}

func commitHasFile(commit *object.Commit, fileName string) (bool, error) {
	stats, err := commit.Stats()
	if err != nil {
		return false, err
	}

	for _, fileStat := range stats {
		if fileStat.Name == fileName {
			return true, nil
		}
	}
	return false, nil
}

const tableWidth = 120

// findOwners uses the information in git log to calculate the owners
// of the files passed in args
// It ultimately prints a table like this
// ****************************************
// |          Ownership for FILE          |
// * **************************************
// | All commits     | Last 50 Commits    |
// * **************************************
// | Touches|Changes |Touches|Changes     |
// * **************************************
// |  user 1| user2  | user3 | user4      |
// ...
func findOwners(cmd *cobra.Command, args []string) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	queries := []*fileLogQuery{
		&fileLogQuery{make(map[string]*ownerTotals), maxRepoDepth, fmt.Sprintf("All (up to %d) commits", maxRepoDepth), 0, 0},
		&fileLogQuery{make(map[string]*ownerTotals), commitsToQuery, fmt.Sprintf("Last %d commits", commitsToQuery), 0, 0},
	}

	for _, file := range args {
		for _, query := range queries {
			commits, err := repo.Log(&git.LogOptions{FileName: &file, Order: git.LogOrderCommitterTime})
			commitsSeen := 0
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			err = commits.ForEach(func(commit *object.Commit) error {
				if commitsSeen >= query.maxCommits {
					return nil
				} else {
					commitsSeen++
				}

				fileInCommit, err := commitHasFile(commit, file)
				if err != nil {
					return err
				}
				if !fileInCommit {
					return nil
				}

				author := commit.Author.Name
				authorTallies := query.ensureOwnerTotals(author)

				commitErr := authorTallies.updateTalliesForFile(file, commit)
				if commitErr != nil {
					return commitErr
				}

				changeCount, err := findChangesForFileInCommit(file, commit)
				if err != nil {
					return err
				}

				// if the file wasn't in the commit, as sometimes happens
				if changeCount > 0 {
					query.totalTouches += 1
					query.totalChanges += changeCount
				}
				return nil
			})

			if err != nil && err != io.EOF {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

		}
		printQueriesInTable(file, queries)
	}
}

// topScorersInTallies returns the top N scores in a list of tallies
// (which must already be sorted) where N is either 3 or the length of the list
func topScorersInTallies(tallies []*ownerTotals) []*ownerTotals {
	// make a new slice and return it, since the slice gets resorted in various ways
	returnTallies := make([]*ownerTotals, 0, ownersToShow)
	for index, tallies := range tallies {
		if index >= ownersToShow {
			return returnTallies
		}

		returnTallies = append(returnTallies, tallies)
	}
	return returnTallies
}

// nthOwner takes in a list of tallies and either returns the owner at that index
// or an empty string
func nthOwner(n int, tallies []*ownerTotals) string {
	if n < len(tallies) {
		return tallies[n].owner
	} else {
		return ""
	}
}

// nthTouches finds the nth item in tallies and returns its touches count
func nthTouches(n int, tallies []*ownerTotals) int {
	if n < len(tallies) {
		return tallies[n].totalTouches
	} else {
		return 0
	}
}

// nthChanges finds the nth item in tallies and returns its changes count
func nthChanges(n int, tallies []*ownerTotals) int {
	if n < len(tallies) {
		return tallies[n].totalChanges
	} else {
		return 0
	}
}

// formatNamePercentage takes a name and a decimal value and returns a string
// in the form of (name (X%))
func formatNamePercentage(name string, percent float32) string {
	if name == "" {
		return ""
	}

	return fmt.Sprintf("%s: %.0f%%", name, (percent * 100))
}

// printQueriesInTable calculates the top owners for each category and prints out the table
// of results
func printQueriesInTable(file string, queries []*fileLogQuery) {
	// at this point, we've gathered all the data about who has committed what
	allTallies := queries[0].ownerTotals()
	// sort by touches
	sort.Slice(allTallies, func(i, j int) bool {
		// want a reverse sort, so function checks greater than
		return allTallies[i].totalTouches > allTallies[j].totalTouches
	})
	topTouchersAll := topScorersInTallies(allTallies)

	sort.Slice(allTallies, func(i, j int) bool {
		// want a reverse sort, so function checks greater than
		return allTallies[i].totalChanges > allTallies[j].totalChanges
	})
	topChangersAll := topScorersInTallies(allTallies)

	recentTallies := queries[1].ownerTotals()
	sort.Slice(recentTallies, func(i, j int) bool {
		// want a reverse sort, so function checks greater than
		return recentTallies[i].totalTouches > recentTallies[j].totalTouches
	})
	topTouchersRecent := topScorersInTallies(recentTallies)

	sort.Slice(recentTallies, func(i, j int) bool {
		// want a reverse sort, so function checks greater than
		return recentTallies[i].totalChanges > recentTallies[j].totalChanges
	})
	topChangersRecent := topScorersInTallies(recentTallies)

	// print out an empty table that will get overwritten later
	printTableDividingLine()
	fmt.Printf("|%-118s|\n", "Ownership for "+file)
	printTableDividingLine()
	fmt.Printf("|%-59s|%-58s|\n", queries[0].description, queries[1].description)
	printTableDividingLine()
	printTableRow("Touches", "Changes", "Touches", "Changes")
	printTableDividingLine()
	for index := 0; index < ownersToShow; index++ {
		value1 := formatNamePercentage(nthOwner(index, topTouchersAll), float32(nthTouches(index, topTouchersAll))/float32(queries[0].totalTouches))
		value2 := formatNamePercentage(nthOwner(index, topChangersAll), float32(nthChanges(index, topChangersAll))/float32(queries[0].totalChanges))
		value3 := formatNamePercentage(nthOwner(index, topTouchersRecent), float32(nthTouches(index, topTouchersRecent))/float32(queries[1].totalTouches))
		value4 := formatNamePercentage(nthOwner(index, topChangersRecent), float32(nthChanges(index, topChangersRecent))/float32(queries[1].totalChanges))
		printTableRow(value1, value2, value3, value4)
	}
	printTableDividingLine()

}

func printTableDividingLine() {
	fmt.Println(strings.Repeat("-", tableWidth))
}

func printTableRow(value1, value2, value3, value4 string) {
	fmt.Printf("|%-29s|%-29s|%-29s|%-28s|\n", value1, value2, value3, value4)
}
