package domain

import (
    "encoding/json"
    "sort"
    "log"
)

type Input struct{
    Index   int
    NumA    int64
    NumB    int64
}

type Output struct{
    Index   int
    NumA    int64
    NumB    int64
    Result  int64
}


// Modeled off of :
// - https://gobyexample.com/worker-pools
//
// We will run several cocurrent instances of the worker
// - job chanel is read only for Input
// - results chanel is write only for Output
func NewWorker(jobs <-chan Input, results chan<- Output) {
	for input := range jobs {
		
		sum := input.NumA + input.NumB

		r := Output{
			Index:      input.Index,
			NumA:		input.NumA,
			NumB:		input.NumB,
			Result:		sum,
		}
		results <- r
	}
}

//For more information on marshelling, see :
//-  http://mattyjwilliams.blogspot.ca/2013/01/using-go-to-unmarshal-json-lists-with.html
// Takes a byte slice and generates an Input slice
func ProcessJsonInput(inputData []byte) []Input {
    collection := []Input{}
    var data map[string][]json.RawMessage
    err := json.Unmarshal(inputData, &data)
    if err != nil {
        log.Println(err)
        return collection
    }
    for _, thing := range data["table"] {
        collection = addInput(thing, collection)

    }
    return collection
}

//The results of the workers' calculations are accumulated in a map. 
//We generate our JSON result with this map. Before we can send 
//it back as a marshalled map of Output structs, we want to sort it by key. 
//The key for the map is the Index
func GenerateJsonOutput(output map[int]Output) ([]byte, error) {
    //The Output is not sorted by index.
    //We sort it by Index prior to returning the response
    sorted := make([]Output, len(output))
    keys := make([]int, len(output))

    for k, _ := range output {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    for _, k := range keys {
        sorted[k] = output[k]
    }
    return json.Marshal(sorted)
}

//Takes a json.RawMessage, converts it to an Input stuct 
//and adds it to our collection
func addInput(thing json.RawMessage, collection []Input) []Input {
    input := Input{}
    err := json.Unmarshal(thing, &input)

    if err != nil {
        log.Println(err)        
    } else {
        if input != *new(Input) {
            collection = append(collection, input)
        }
    }

    return collection
}