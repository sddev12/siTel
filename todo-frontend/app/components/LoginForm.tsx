import React, { useState } from "react";
import TextInput from "./TextInput";
import Button from "./Button";
import Link from "next/link";

export default function LoginForm() {
  const [username, setUsername] = useState<string>("");
  const [usernameError, setUsernameError] = useState<boolean>(false);
  const handleLoginClick = () => {
    console.log("Login Button Clicked");
  };

  return (
    <>
      <TextInput
        type={"text"}
        placeholder={"username"}
        username={username}
        setUsername={setUsername}
        usernameError={usernameError}
        setUsernameError={setUsernameError}
      />
      {usernameError && <p className="text-red-600 m-2">Enter a username</p>}
      <Button
        text={"Login"}
        username={username}
        setUsernameError={setUsernameError}
        handleClick={handleLoginClick}
      />
      <Link href="/register">Register</Link>
    </>
  );
}
