export default function TextInput({
  type,
  placeholder,
  username,
  setUsername,
  usernameError,
  setUsernameError,
}: any) {
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(event?.target.value);

    if (username) {
      if (usernameError) {
        setUsernameError(false);
      }
    }
  };

  return (
    <>
      <div>
        <input
          className="text-black h-12 w-80 p-2 m-2 rounded"
          type={type}
          placeholder={placeholder}
          name=""
          id=""
          onChange={handleChange}
        />
      </div>
    </>
  );
}
