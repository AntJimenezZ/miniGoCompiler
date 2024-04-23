import React, { useState } from "react";
import { Button, Textarea } from "@nextui-org/react";


export default function App() {
  const [text, setText] = useState("");
  const [responseText, setResponseText] = useState("");
  const [backupCode, setBackupCode] = useState("");
  const [isReadOnly, setIsReadOnly] = useState(false);

  const handleClickCompile = async () => {
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
      
      const newBackupCode = text;
      setBackupCode(newBackupCode);
      console.log("Backup code: ", newBackupCode);
      setIsReadOnly(true);

      setText(text.split('\n').map((line, index) => `${index + 1}- ${line}`).join('\n'));
      
      
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const handleClickEdit = async () => {
    
    console.log("Backup code: ", backupCode);
    setText(backupCode);
    setIsReadOnly(false);
  }

  return (
    <div>
      <div className="mb-8 mt-8">
        <h1 className="mb-8 ite">Enter your code below</h1>
        <Textarea
          label="Code:"
          variant="flat"
          disableAnimation
          //disableAutosize
          isReadOnly={isReadOnly}
          value={text}
          onChange={(e) => setText(e.target.value)}
          classNames={{
            base: "w-full max-w-7xl",
            input: "resize-y min-h-[270px] font-mono",
          }}
        />
        <h1>
          <Button color="success" variant="shadow" size="lg" radius="sm" isDisabled ={isReadOnly} onPress={handleClickCompile} style={{marginRight: '20px'}}>
            Compile
          </Button>
          <Button color="warning" variant="shadow" size="lg" radius="sm" isDisabled ={!isReadOnly} onPress={(handleClickEdit)}>
            Edit
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
