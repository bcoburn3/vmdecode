package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var noArgOps = "01,03,04,0c,0d,0e,0f,10,11,12,13,14,15,16,17,18,19,1a,1b,1c,1d,1e,1f,20,21,34,38"

var oneArgOps = "00,02,05,06,07,08,09,0a,0b,22,37"

var codeNameMap = map[string]string{
	"00": "LIT",
	"01": "BACK",
	"02": "REL",
	"03": "SWAP",
	"04": "POP",
	"05": "IMM",
	"06": "JMP",
	"07": "JSR",
	"08": "BZ",
	"09": "BNZ",
	"0a": "ENT",
	"0b": "ADJ",
	"0c": "RET",
	"0d": "LI",
	"0e": "LC",
	"0f": "SI",
	"10": "SC",
	"11": "PUSH",
	"12": "OR",
	"13": "XOR",
	"14": "AND",
	"15": "EQ",
	"16": "NE",
	"17": "LT",
	"18": "GT",
	"19": "LE",
	"1a": "GE",
	"1b": "SHL",
	"1c": "SHR",
	"1d": "ADD",
	"1e": "SUB",
	"1f": "MUL",
	"20": "DIV",
	"21": "MOD",
	"22": "INT",
	"34": "PUSHARG",
	"37": "JSRP",
	"38": "RETP",
}


func main() {
	hs := getHexString(os.Args[1])
	hsReader := strings.NewReader(hs)
	tempBytes := make([]byte, 2)
	n, _ := hsReader.Read(tempBytes)
	for string(tempBytes) != "0a" {
		if n == 0 {
			fmt.Println("could not find start of VM function, bytecode 0x0a")
			panic("no vm func")
		}
		n, _ = hsReader.Read(tempBytes)
	}

	argBytes := make([]byte, 8) //ENT bytecode takes one argument
	n, _ = hsReader.Read(argBytes)
	if n < 8 {
		fmt.Println("missing argument to function entrypoint bytecode, did you copy the whole thing?")
		panic("no vm arg")
	}
	fmt.Println("ENT " + string(argBytes))
	
	line, _ := readAssemOp(hsReader)
	for line != "" {
		fmt.Println(line)
		line, _ = readAssemOp(hsReader)
	}
	fmt.Println(line)
}

func getHexString(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tempSlice := []string{}
	for scanner.Scan() {
		tempSlice = append(tempSlice, scanner.Text()[5:53])
	}
	res := strings.Join(tempSlice, "")
	res = strings.Replace(res, " ", "", -1)
	res = strings.Replace(res, "\t", "", -1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func readAssemOp(hsReader *strings.Reader) (string, error) {
	tempBytes := make([]byte, 2)
	n, err := hsReader.Read(tempBytes)
	if n == 0 {
		return "", err
	}
	opcode := string(tempBytes)
	if strings.Contains(noArgOps, opcode) {
		return codeNameMap[opcode], err
	} else if strings.Contains(oneArgOps, opcode) {
		argBytes := make([]byte, 8)
		_, err := hsReader.Read(argBytes)
		return codeNameMap[opcode] + " " + string(argBytes), err
	} else {
		return "UNDEF", err
	}
}
	




