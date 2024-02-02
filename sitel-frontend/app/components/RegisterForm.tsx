import { useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";
import Link from "next/link";
import TextInput from "./TextInput";
import Button from "./Button";

export default function RegisterForm() {
  const router = useRouter();
  const [username, setUsername] = useState<string>("");
  const [usernameError, setUsernameError] = useState<boolean>(false);

  const handleRegisterClick = async () => {
    console.log("Register Button Clicked");
    const iamServiceHost = process.env.NEXT_PUBLIC_IAM_SERVICE_HOST;
    const iamServicePort = process.env.NEXT_PUBLIC_IAM_SERVICE_PORT;
    const iamServiceUrl = `http://${iamServiceHost}:${iamServicePort}/register`;

    console.log("making request to iam /register");
    const registerResponse = await fetch(iamServiceUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
      }),
    });

    if (registerResponse.ok) {
      const responseData = await registerResponse.json();
      console.log(responseData.sessionId);
      Cookies.set("sessionId", responseData.sessionId);
      router.push("/login");
    }
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
