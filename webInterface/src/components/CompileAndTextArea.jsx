import React, { useState } from "react";
import { Button, Textarea } from "@nextui-org/react";

export default function App() {
  const [text, setText] = useState("");
  const [responseText, setResponseText] = useState("");

  const handleClick = async () => {
    try {
      const response = await fetch("http://localhost:8080/json", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: text }),
      });

      const data = await response.json();
      setResponseText(data.text);
      console.log(data.text);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <div className="mb-8 mt-8">
        <h1 className="mb-8 ite">Enter your code below</h1>
        <Textarea
          label="Code:"
          variant="flat"
          placeholder="Example:
1 package main;
2 import fmt;
3 func main(){ 
4       fmt.Println('Hello, World!'); 
5 }"
          disableAnimation
          //disableAutosize
          value={text}
          onChange={(e) => setText(e.target.value)}
          classNames={{
            base: "w-full max-w-7xl",
            input: "resize-y min-h-[270px] font-mono",
          }}
        />
        <h1>
          <Button color="success" variant="solid" size="lg" radius="sm" onPress={handleClick}>
            Compile
          </Button>
        </h1>
      </div>

      <div className="mb-96">
        <h1 className="mb-8">All errors will be displayed here</h1>
        <Textarea
          label="Errors:"
          variant="flat"
          isReadOnly={true}
          placeholder="When you click the 'Compile' button, the code will be sent to the server, compiled, and the result will be displayed here."
          disableAnimation
          value={responseText}
          classNames={{
            base: "w-full max-w-7xl text-black",
            input: "resize-y min-h-[270px] font-mono",
          }}
        />
      </div>
    </div>
  );
}
