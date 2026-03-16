# KATH DOCUMENTATION

## how kath works

kath code operates on a 4096 long array of integers, referred to in code as "memory". Kath code manipulates that array by keeping track of a pointer to an array position and manipulating the data at that position.

## syntax

kaths syntax is extremely minimal. A kath program consists of multiple lines, each line consisting of an instruction and an argument. Note that arguments do not always influence what the instruction does, but they still have to be provided. It is recommended to always set these "irrelevant" argument to 0

All kath programs need an entrypoint, that being the procedure ```proc 0```. If this procedure doesnt exist, the code will fail to execute. 
## keywords
- pnt:
    - sets the pointer to its argument. Fails if the provided argument is an invalid memory position.
- shift:
    - moves the pointer by adding the position of the argument to its current position. Fails if the new position is an invalid memory position.
- add:
    - adds the value of the memory position specified in the argument. Fails if the provided argument is an invalid memory position.
- addn:
    - adds the argument to the memory position that the pointer points to. 
- sub:
    - subtracts the value of the memory position specified in the argument. Fails if the provided argument is an invalid memory position.
- subn:
    - adds the argument to the memory position that the pointer points to. 
- copy:
    - copies the value of the memory position specified by the argument to the current pointer position. Fails if the provided argument is an invalid memory position. 
- in:
    - accepts a 64 bit integer from the standard input and puts it in the current memory position
- out:
    - prints the value of the current memory position to the standard output
- loop/unloop:
    - each loop statement needs a corresponding unloop statement
    - loop starts the loop if the value of the pointer position doesnt match the argument, otherwise skips to its corresponding unloop
    - unloop exits the loop if the value of the pointer does match its argument, otherwise skips to its corresponding loop statement, restarting the loop
- proc:
    - creates a new procedure. Requires an end statement at the end of the procedure. The argument acts as an identifier for the procedure, letting you specify which procedure to call from an eval statement. Proc 0 is the entrypoint. 
- end: 
    - signals the end of a procedure. Argument is irrelevant.
- eval:
    - evaluates the procedure with an identifier matching the provided argument. Fails if that procedure doesnt exist.
    - While possible to recursively call proc 0, it probably is not a good idea
