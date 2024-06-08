### JSON Parser

The JSON Parser is a tool designed to parse and manipulate JSON (JavaScript Object Notation) data. It allows you to easily extract, modify, and analyze JSON data structures using a simple and clean interface.

#### What I have Learnt by Building This Project
- Understood parsing techniques
- Handled simple data formats
- Developed a fully featured compiler using Go Programming Language

#### Stages of Parsing

Parsing involves several stages

1. **Lexical Analysis**  (Word Breakdown):
    - **Tokenization**: Breaking down the code into tokens (words, symbols, other meaningful elements).
    - **Classification**: Grouping these parts into types like keywords (if, for), variables (x, y), and operators (+, -).

2. **Syntax Analysis** (Sentence Structure):
    - **Parsing**: Building a syntax tree based on the tokens using a set of grammatical rules.
    - **Error Detection**: Finding and reporting syntax errors in the code's structure

3. **Semantic Analysis**  (Meaning Check):
    - **Type Checking**: Making sure the operations make sense, like adding integers, not words.
    - **Scope Resolution**: Checking if all variables are properly defined and used.
   
4. **Intermediate Code Generation** (Middle Step):
    - **AST Transformation**: Turning the structure into a simpler form that’s easier to work with.
    - **Basic Code Emission**: Creating an early version of the final code.

5. **Optimization** (Improvement):
    - **Intermediate Code Optimization**: Making the middle step code better and faster.
    - **Machine-Independent Optimization**: Improving the code in a way that works on any machine.

6. **Code Generation** (Final Step):
    - **Target Code Emission**: Converting the improved middle step code into machine language.
    - **Instruction Selection and Scheduling**: Choosing the best instructions and ordering them efficiently.

7. **Code Optimization** (Fine-Tuning):
    - **Machine-Specific Optimization**: Making the machine code better for the specific computer it will run on.
    - **Register Allocation**: Efficiently assigning variables to the computer’s storage.

8. **Code Linking and Loading** (Finishing Up):
    - **Linking**: Combining all parts of the code and libraries into one executable file.
    - **Loading**: Preparing this file to run on the computer, including setting up memory.
