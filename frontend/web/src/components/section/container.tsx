import { ReactNode } from "react"

type Props = {
    children: ReactNode
}
export function Container({children} : Props){
    return (
        <div className="py-12 flex flex-wrap items-center w-full max-w-[1246px] px-4 mx-auto justify-between">
            {children}
        </div>
    )
}