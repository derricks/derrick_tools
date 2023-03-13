/*
Copyright © 2022 Derrick Schneider derrick.schneider@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

// Tools for making "cross queries" on objects used by memory quiz.
// For instance, if a country struct has a number, a name, and a currency,
// we should be able to automatically construct memory quizzes that ask about name from size,
// size from name, currency from name, currency from size, and so on.
// example struct annotation
// type country {
//   name string `crossquery:"all"`
//   currency string `crossquery:"guess"` -- because so many countries use the same currency name, this can't be a given "What country has a currency of franc"
//   capital string `crossquery:""`
// }
// currently only supports strings and ints

func constructCrossQuery(entityType string, entity interface{}) promptAndResponse {
	// the fields we could use as the "given" in the prompt. e.g., the country's name is Algeria
	// these are fields annotated with crossquery:all or crossquery:given
	givens := []reflect.StructField{}
	// the fields we could use as things to guess in the prompt. e.g., the currency
	// these are fields annotated with crossquery:all or crossquery:given
	guesses := []reflect.StructField{}

	reflectStruct := reflect.TypeOf(entity)
	for i := 0; i < reflectStruct.NumField(); i++ {
		field := reflectStruct.Field(i)
		if crossQuery, ok := field.Tag.Lookup("crossquery"); ok {
			if crossQuery == "" || crossQuery == "all" {
				givens = append(givens, field)
				guesses = append(guesses, field)
			} else if crossQuery == "given" {
				givens = append(givens, field)
			} else if crossQuery == "guess" {
				guesses = append(guesses, field)
			} else {
				// this is effectively a syntax error, so kill the program
				panic(fmt.Sprintf("Invalid value for crossquery: %s", crossQuery))
			}
		}
	}

	reflectEntity := reflect.ValueOf(entity)

	//get a given and figure out a non-equal guess
	given := givens[rand.Intn(len(givens))]
	guess := guesses[rand.Intn(len(guesses))]
	for given.Name == guess.Name {
		guess = guesses[rand.Intn(len(guesses))]
	}

	givenValueRaw := reflectEntity.FieldByName(given.Name)
	guessValueRaw := reflectEntity.FieldByName(guess.Name)

	// convert to strings if need be
	givenValue := reflectValueToString(givenValueRaw)
	guessValue := reflectValueToString(guessValueRaw)
	return promptAndResponse{fmt.Sprintf("What is the %s of the %s with %s of %v?", guess.Name, entityType, given.Name, givenValue), guessValue}
}

func reflectValueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.Itoa(int(v.Int()))
	case reflect.Slice:
		// grab an item at random from the slice. Note that what you get back from Index is a value
		return reflectValueToString(v.Index(rand.Intn(v.Len())))
	default:
		return v.String()
	}
}
