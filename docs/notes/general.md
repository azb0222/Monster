# GENERAL 

- **dsl**: domain-specific language; pidgins for a specific task ie. YAML, JSON, VimScript, HTML, etc.


## Interpreter Parts 

- "interpreter parts are a pipeline: transforming data into more organized formats for the next step of the pipeline"

#### Front end 
1) **scanning/lexical analysis**: take in linear stream of characters, outputs a list of tokens
   - token: can be single character or several characters long 
   - scanner discards whitespace/meaningless characters 
2) **parsing**: using a **grammar**, parser takes list of tokens, outputs a parse tree
   - parser can report **syntax errors**
3) **static analysis**: 
    - perform **binding/resolution**: associate **scope** with each **identifier**
    - perform **type checking** for statically typed languages
    - store data in **symbol table**
#### Middle end 
4) **IR**: code may be stored in IR (intermediate representation); interface between frontend and backend 
<TODO: left off at page 14> 