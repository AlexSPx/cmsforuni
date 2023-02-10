import LinkButton from '@/Components/LinkButton'
import Link from 'next/link'
import React from 'react'

export default function Navbar() {
  return (
    <div className="bg-white border-gray-200 px-4 lg:px-6 py-2 dark:bg-gray-800">
      <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl">
        <Link href="/" className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">+open</Link>

        <div className="flex items-center lg:order-2">
            <LinkButton label='Sign Up' href='/signup'/>
          </div>
      </div>

    </div>
  )
}
