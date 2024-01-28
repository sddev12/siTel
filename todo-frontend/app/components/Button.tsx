import { error } from "console";

export default function Button({
  text,
  handleClick,
  username,
  setUsernameError,
}: any) {
  const errorProcessor = (): undefined => {
    if (username == "") {
      setUsernameError(true);
      return;
    }

    handleClick();
  };
  return (
    <>
      <div
        id="submit-button"
        className="flex justify-center items-center p-2 m-5 rounded bg-white text-black font-bold w-40 hover:bg-black hover:text-white hover:border"
        onClick={errorProcessor}
      >
        <p>{text}</p>
      </div>
    </>
  );
}
