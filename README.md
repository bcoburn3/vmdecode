# vmdecode
Convert stockfighter.io Jailbreak memory into VM assembly

To use:
  1. Download vmdecode.go and level1.txt
  2. go install /path/to/vmdecode.go
  3. $GOPATH/bin/vmdecode /path/to/level1.txt

You should see output something like this:

ENT 00000000

IMM 0000008d

PUSHARG

INT 00000002

ADJ 00000001

IMM 00000000

RET

RET

LIT 
