import LoginInput from "./LoginInput";

export default function Home() {
  return (
    <main className="flex flex-col w-screen h-screen items-center">
      <div className="flex w-full h-2/3 bg-gray-100 items-center p-24">
        <h2 className="text-5xl font-extrabold leading-none tracking-tight font-montserrat md:text-6xl">Notes for <br /> university and <br /> blog posts</h2>
        <section className="flex justify-center w-2/3">
        <div className="w-1/2 h-full">
          <LoginInput />
        </div>
        </section>
      </div>
    </main>    
  )
}
