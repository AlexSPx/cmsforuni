import { Dispatch, SetStateAction } from "react"

type InputFieldType = {
    label: string,
    type: "email" | "text" | "password",
    value: string,
    setValue: Dispatch<SetStateAction<string>>
} 

export default function InputField({label, type, value, setValue}: InputFieldType) {
    return (
    <div className="relative flex items-center mt-3 w-full">
        <input type={type} className="block w-full py-3 text-gray-700 bg-white border rounded-lg px-11 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 dark:focus:border-blue-300 focus:ring-blue-300 focus:outline-none focus:ring focus:ring-opacity-40" 
            placeholder={label} 
            value={value} 
            onChange={(e) => {
                setValue(e.target.value)
            }}
        />
    </div> 
    )
}