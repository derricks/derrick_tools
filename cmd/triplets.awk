# A script to parse spelling_bee.go and look for triplets
# of letters that form each member of the family. The goal is to
# go from the passive strategy of seeing one word in a word family
# and knowing all the others to a more active strategy where I
# memorize what families _might_ be found if a given triplet
# is on the board.
# e.g., AIR means I should look for FAIR, etc., TIARA, etc. and others

# assumes that there's a global variable family_members with
# the list of family members
# sets a global variable triplets that's each of the normalized
# triplets that can be found in a word

# requires gawk for asort

func calculate_triplets() {
    delete triplets
    for (word_idx in family_members) {
        word = family_members[word_idx]

        # for each word, make three loops starting with the first letter
        # in the inner loop, construct a triplet of letters
        # then normalize (so GNI and ING, etc. become GIN)
        for (i = 1; i <= length(word); i++) {
            first_letter = substr(word, i, 1)
            for (j = 1; j <= length(word); j++) {
                second_letter = substr(word, j, 1)
                if (second_letter == first_letter) {
                    continue
                }

                for (k = 1; k <= length(word); k++) {
                    third_letter = substr(word, k, 1)
                    if (third_letter == second_letter || third_letter == first_letter) {
                        continue
                    }

                    triplet[1] = first_letter
                    triplet[2] = second_letter
                    triplet[3] = third_letter

                    asort(triplet)
                    cur_triplet = (triplet[1] triplet[2] triplet[3])
                    if (triplets[cur_triplet] == 0) {
                        triplets[cur_triplet] = 1
                    }
                }
            }
        }
    }
}

# figures out the spelling bee score for a given family
# assumes a global variable family_members which is an array
# of the members in a family
func score_family() {
    family_score = 0
    for (word_idx in family_members) {
        word = family_members[word_idx]
        if (length(word) == 4) {
            family_score += 1
        } else {
            family_score += length(word)
        }
    }
    return family_score
}

# get just the slices of strings
/\{"/{
   # some clean-up
   # get rid of the go syntax stuff
   gsub("[\t {}\"]", "", $0)
   # get rid of the trailing comma
   $0 = substr($0, 1, length($0) - 1)
   family_count = split($0, family_members, ",")
   calculate_triplets()
   family_score = score_family()
   for (current_triplet in triplets) {
      triplet_to_score[current_triplet] += family_score
      triplet_to_count[current_triplet] += 1
      triplet_to_families[current_triplet] = (triplet_to_families[current_triplet] "," family_members[1])
   }
}

END {
    for (current_triplet in triplet_to_score) {
        print current_triplet, triplet_to_score[current_triplet], triplet_to_count[current_triplet], triplet_to_families[current_triplet]
    }
}