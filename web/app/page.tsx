'use client'

import { VerifyToken } from "@/utils/verifyToken";
import Image from "next/image";
import { useRouter } from 'next/navigation'
import { useEffect, useState } from "react";
import Cookies from 'js-cookie';
import { axiosInstante } from "@/utils/axiosInstante";

export default function Home() {
  const router  = useRouter();

  const [ email, setEmail ] = useState("");
  const [ password, setPassword ] = useState("");
  const [ msg, setMessage ] = useState(null);
  const [ isLoading, setIsLoading ] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await VerifyToken(Cookies.get("kukis") as string)
        if (data?.code === 200) {
          router.push("/dashboard")
        }
      } catch (e: any) {
        router.push("/")
      }
    }

    fetchData();

  }, [router]);

  const handleSubmit = async (e: any) => {
    e.prevenDefault()
    setIsLoading(true)

    try {
      const response = await axiosInstante.post("/api/user/login", {
        email,
        password
      })

      Cookies.set("kukis", response.data.data)
      setIsLoading(false)
      router.push("/dashboard")
      
    } catch (error: any) {
      setMessage(error?.response?.data.error.Message);
      setIsLoading(false)

    }
  }
  return (
    <div className="pt-52 pb-16">
      <div className="flex flex-wrap items-center justify-center">
        <div className="px-4">
          <div className="grid grid-cols-1">
            <Image
              src={"/next.svg"}
              alt="Image Profile"
              width={300}
              height={300}
              className="pb-5"
            />
            <form onSubmit={handleSubmit} action="post" className="text-sm">
              <div className="mt-5 my-2">
                <label htmlFor="email" className="text-balance text-sm font-bold font-jetBrains  text-semiDark">Email <span className="text-red-500">*</span>
                </label>
                <input
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  type="email" 
                  className="text-darkColor font-bold mt-1 placeholder:text-black/30 w-full bg-transparent pt-3 pl-2 pb-2 border border-darkColor/60 focus:outline-none rounded-lg"
                  placeholder="email@email.com"
                  id="email"
                  required
                />
              </div>
              <label htmlFor="password" className="text-balance text-sm font-bold font-jetBrains  text-semiDark">Password <span className="text-red-500">*</span>
              </label>
              <div className="relative mb-6">
              <input
                  value={password}
                  onChange={(e) => setEmail(e.target.value)}
                  type="password" 
                  className="text-darkColor font-bold mt-1 placeholder:text-black/30 w-full bg-transparent pt-3 pl-2 pb-2 border border-darkColor/60 focus:outline-none rounded-lg"
                  placeholder="*******"
                  id="password"
                  required
                />
              </div>
              <button
                type="submit"
                className="w-full text-lightColor bg-darkColor px-5 py-2 transition-full-200 rounded-lg"
              >
                {isLoading ? "Loading..." : "Login"}
              </button>

              <p className="text-center test-sm text-red-500 capitalize">
                {msg ? "Your email or password is wrong" : null}
              </p>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}