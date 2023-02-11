/* eslint-disable @next/next/no-head-element */
import { Ubuntu } from '@next/font/google'
import Navbar from './navbar';
import './globals.css';

const font = Ubuntu({weight: "400"})

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <head></head>
      <body className={`flex flex-col w-screen h-screen`}>
        <Navbar />
        {children}
      </body>
    </html>
  );
}
