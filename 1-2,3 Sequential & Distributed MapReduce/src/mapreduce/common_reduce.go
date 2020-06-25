package mapreduce

import (
	"os"
	"encoding/json"
)

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Use checkError to handle errors.
	
	key2values := make(map[string][]string)
	for i:=0; i<nMap; i++{
		f, err := os.Open(reduceName(jobName,i,reduceTaskNumber))
		checkError(err)
		defer f.Close()
		decoder := json.NewDecoder(f)
		var kv KeyValue
		for decoder.Decode(&kv) == nil {
			if values,found := key2values[kv.Key]; found{
				key2values[kv.Key] = append(values,kv.Value)
			} else {
				key2values[kv.Key] = []string{kv.Value}
			}
		}
	}
	mergeFile, err := os.OpenFile(mergeName(jobName, reduceTaskNumber), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError(err)
	enc := json.NewEncoder(mergeFile)
	for key, values := range key2values {
		enc.Encode(KeyValue{key, reduceF(key, values)})
	}
	mergeFile.Close()
}
