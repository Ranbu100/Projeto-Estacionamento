
type Props = {
    nome: string;
}

export function ItemMenu({nome} : Props){
    return (
        <button className="flex items-center gap-1 sm:gap-3">
            <span className="text-white">{nome}</span>
        </button>
    )
}