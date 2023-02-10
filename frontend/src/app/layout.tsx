/* eslint-disable @next/next/no-head-element */
import { Inter } from '@next/font/google'
import Navbar from './navbar';
import './globals.css';

const inter = Inter({ subsets: ['latin'] })

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <head></head>
      <body className={`${inter.className} flex flex-col w-screen h-screen`}>
        <Navbar />
        {children}
      </body>
    </html>
  );
}
