"use client"

import Button from "@/Components/Button"
import InputField from "@/Components/InputField"
import { useState } from "react"

export default function LoginInput() {

    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")

  return (
    <section className="flex flex-col w-full h-full items-center justify-center">
         <h1 className="mt-4 text-2xl font-semibold tracking-wide text-center text-gray-800 capitalize md:text-3xl dark:text-white">
            Login to your account
        </h1>
        <div className="my-2 w-full">
            <InputField label="Username" value={username} setValue={setUsername} type="text"/>
            <InputField label="Password" value={password} setValue={setPassword} type="password"/>
        </div>
        <Button label="Login" func={() => {}}/>
    </section>
  )
}
