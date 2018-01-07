package main

import (
    "net/http"
    "os/exec"
    "fmt"
//    "log"
)

//workload
var amt_of_data string = "34M"//os.Getenv("amt_of_data")
var max_prime string = "4000"//os.Getenv("max_prime")
var num_threads string = "2"//os.Getenv("num_threads")
var file_size string = "300K"//os.Getenv("file_size")

//network parameter
const num_reply_bytes int = 20
var talks_to string = "golangdemob:8080/start"

func generate_memory_workload(amount_of_data string){
     exec.Command("sysbench","memory", "--memory-total-size="+amount_of_data ,"run").Output()
}

func generate_cpu_workload(max_prime string,num_threads string){
     exec.Command("sysbench", "--test=cpu", "--cpu-max-prime="+max_prime, "--num-threads="+num_threads, "run").Output()
}

func generate_read_and_write_file_workloads(file_size string){
     exec.Command("sysbench", "--test=fileio", "--file-total-size="+file_size ,"prepare").Output()
     exec.Command("sysbench", "--test=fileio", "--file-total-size="+file_size , "--file-test-mode=rndrw", "--init-rng=on", "--max-time=300", "--max-requests=0", "run").Output()
     exec.Command("sysbench", "--test=fileio", "--file-total-size="+file_size ,"cleanup").Output()
}

func workload_gen(amt_of_data string, max_prime string, num_threads string, file_size string){
     generate_memory_workload(amt_of_data)
     generate_cpu_workload(max_prime,num_threads)
     generate_read_and_write_file_workloads(file_size) // more than available ram
}

func send_message(comp string){
     resp, err := http.Get("http://"+comp)
     if err != nil {
    // handle error
       fmt.Printf("get error %s\n",comp);
      }else{
     	fmt.Printf("response : %s \n",resp);
      }
}

func start(w http.ResponseWriter, r *http.Request) {

     workload_gen(amt_of_data,max_prime,num_threads,file_size)
     send_message(talks_to)
     w.WriteHeader(http.StatusOK)
     w.Write(make([]byte, num_reply_bytes))
}

func main() {
    http.HandleFunc("/start", start)
    http.ListenAndServe(":8080", nil)
}
