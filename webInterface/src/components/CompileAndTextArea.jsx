import React,{useState} from "react";
import {Button, Textarea} from "@nextui-org/react";

export default function App() {


  const [text, setText] = useState('');
  const [responseText, setResponseText] = useState(''); 
  const [errorText, setErrorText] = useState('');

  const handleClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/json', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ text: text })
      });

      const data = await response.json();
      setResponseText(data.text);
      console.log(data.text);
    } catch (error) {
      console.error('Error:', error);
    }
  };


  return (
    <div className="flex ml-80">
  <div>
    <h1 className="mb-8">
      Write your code here
    </h1>

    <Textarea
      label="Code:"
      variant="flat"
      placeholder="Example:
package main;
import fmt;
func main(){ 
    fmt.Println('Hello, World!'); 
}"
      disableAnimation
      //disableAutosize
      value={text}
      onChange={(e) => setText(e.target.value)}
      classNames={{
        base: "w-full max-w-xl",
        input: "resize-y min-h-[270px]",
      }}
    />
    <h1>
      <Button color="success" variant="shadow" onPress={handleClick}>
        Compile
      </Button>
    </h1>
  </div>


  <div className="ml-72">
    <h1 className="mb-8">
      All errors will be displayed here
    </h1>
    <Textarea
      label="Errors:"
      variant="flat"
      isReadOnly = {true}
      placeholder="When you click the 'Compile' button,
the code will be sent to the server, compiled, and the result will be displayed here."
      disableAnimation
      value={responseText}
      classNames={{
        base: "w-full max-w-xl text-black",
        input: "resize-y min-h-[270px]",
      }}
    />
  </div>
</div>
    
  );
}
