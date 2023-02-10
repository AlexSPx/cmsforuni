"use client"

import InputField from "@/Components/InputField";
import { useState } from "react";


export default function SignupForm() {
    const [username, setUsername] = useState("")
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")


  return (
    <section className="bg-white dark:bg-gray-900">
    <div className="container flex items-center justify-center px-6 mx-auto">
        <form className="w-full max-w-md">    
            <div className="flex items-center justify-center mt-6">
                <a href="#" className="w-1/3 pb-4  text-center text-gray-800 capitalize border-b-2 border-blue-500 dark:border-blue-400 dark:text-white">
                    sign up
                </a>
            </div>

            <InputField label="Username" type="text" value={username} setValue={setUsername} />

            <label className="flex items-center px-3 py-3 mx-auto mt-6 text-center bg-white border-2 border-dashed rounded-lg cursor-pointer dark:border-gray-600 dark:bg-gray-900">
                <h2 className="mx-3 text-gray-400">Profile Photo</h2>
                <input id="dropzone-file" type="file" className="hidden" />
            </label>

            <InputField label="Email" type="email" value={email} setValue={setEmail} />
            <InputField label="Name" type="text" value={name} setValue={setEmail} />
            <InputField label="Password" type="password" value={password} setValue={setPassword} />
            <InputField label="Confirm Password" type="password" value={confirmPassword} setValue={setConfirmPassword} />

            <div className="mt-6">
                <button className="w-full px-6 py-3 text-sm tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-500 rounded-lg hover:bg-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-50">
                    Sign Up
                </button>
            </div>
        </form>
    </div>
</section>
  )
}
