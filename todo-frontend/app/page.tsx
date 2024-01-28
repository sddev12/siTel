"use client";
import { useRouter } from "next/navigation";
import Button from "./components/Button";

export default function Home() {
  const router = useRouter();

  const handleLoginClick = () => {
    router.push("/login");
  };

  const handleRegisterClick = () => {
    router.push("/register");
  };
  return (
    <>
      <p className="text-6xl m-5">Todo</p>
      <div className="flex">
        <Button text="Login" handleClick={handleLoginClick}></Button>
        <Button text="Register" handleClick={handleRegisterClick}></Button>
      </div>
    </>
  );
}
