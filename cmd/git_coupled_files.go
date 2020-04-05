package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"sort"
)

var couplingsToShow int
var coupledFilesCmd = &cobra.Command{
	Use:   "coupled-files",
	Short: "For the given files, determine the files they are most often checked in with",
	Long: `Files often have hidden dependencies in a repo that aren't clear to a new contributor.

  This command looks through the entire commit history of the given file and finds which files were
  committed alongside it each time. It then prints out up to the top 10 files committed with it.
  `,
	Args: cobra.MinimumNArgs(1),
	Run:  findCoupledFiles,
}

func init() {
	coupledFilesCmd.Flags().IntVarP(&couplingsToShow, "count", "c", 10, "The number of top couplings to show")
	gitCmd.AddCommand(coupledFilesCmd)
}

type coupledFile struct {
	name         string
	coupledCount int
}

type coupledFileTracker struct {
	file         string
	coupledFiles map[string]*coupledFile
}

// coupleFile adds a coupling count for file
func (tracker *coupledFileTracker) coupleFile(file string) {
	coupling, exists := tracker.coupledFiles[file]
	if exists {
		coupling.coupledCount += 1
	} else {
		tracker.coupledFiles[file] = &coupledFile{file, 1}
	}
}

// coupledFiles returns just the slice of coupledFile structs
// without the internal details of the mapping
func (tracker *coupledFileTracker) getCoupledFiles() []*coupledFile {
	couplings := make([]*coupledFile, 0, len(tracker.coupledFiles))
	for _, coupling := range tracker.coupledFiles {
		couplings = append(couplings, coupling)
	}
	return couplings
}

func newTracker(file string) *coupledFileTracker {
	return &coupledFileTracker{file, make(map[string]*coupledFile)}
}

func findCoupledFiles(command *cobra.Command, args []string) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, file := range args {
		tracker := newTracker(file)
		totalCommits := 0
		// for each file, look through all the commits and couple every file
		// that was checked in with the given file

		commits, err := repo.Log(&git.LogOptions{FileName: &file, Order: git.LogOrderCommitterTime})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		err = commits.ForEach(func(commit *object.Commit) error {
			// is file even in the commit? if not, we can continue
			fileInCommit, err := commitHasFile(commit, file)
			if err != nil {
				return err
			}
			if !fileInCommit {
				return nil
			}

			totalCommits += 1

			// get all the other files
			stats, err := commit.Stats()
			if err != nil {
				return err
			}

			for _, fileStat := range stats {
				// don't include this file!
				if fileStat.Name != file {
					tracker.coupleFile(fileStat.Name)
				}
			}
			return nil
		})

		coupledFiles := tracker.getCoupledFiles()
		sort.Slice(coupledFiles, func(i, j int) bool {
			// we want the reverse sort
			return coupledFiles[i].coupledCount > coupledFiles[j].coupledCount
		})

		fmt.Printf("Top %d coupled files\n", couplingsToShow)
		for index := 0; index < couplingsToShow; index++ {
			if len(coupledFiles) <= index {
				break
			}

			fmt.Printf("%s: %2.0f%%\n", coupledFiles[index].name, 100.0*float32(coupledFiles[index].coupledCount)/float32(totalCommits))
		}

	}

}
