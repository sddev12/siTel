import { useState } from "react";
import Link from "next/link";
import TextInput from "./TextInput";
import Button from "./Button";

export default function RegisterForm() {
  const [username, setUsername] = useState<string>("");
  const [usernameError, setUsernameError] = useState<boolean>(false);
  const handleRegisterClick = () => {
    console.log("Register Button Clicked");
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
        text={"Register"}
        username={username}
        setUsernameError={setUsernameError}
        handleClick={handleRegisterClick}
      />
      <Link href="/login">Login</Link>
    </>
  );
}
