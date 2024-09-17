'use client'
import {useRouter} from "next/navigation"
export function Login(){
    const Router = useRouter()
    return (

        <button 
          type="button" 
          className="flex items-center font-semibold text-blue-900 h-auto w-auto p-0 lg:p-2 bg-white border"
          onClick={() => Router.push('/login')}
        >
            <span className="h-auto w-auto p-0.5">Logar</span>
        </button>
    )
}