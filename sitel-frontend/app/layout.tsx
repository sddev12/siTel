import "./globals.css";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <div className="w-screen h-screen flex flex-col justify-center items-center">
          {children}
        </div>
      </body>
    </html>
  );
}
