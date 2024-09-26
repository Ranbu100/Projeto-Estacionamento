import { EyeOpenIcon, EyeClosedIcon } from "@radix-ui/react-icons";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

interface Props {
    isVisible: boolean;
    onChangeVisibility: () => void;
    label: string;
    value?: string;
    onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void; // Especificar o tipo do evento
    id: string;
}

export function Senhainput({ isVisible, onChangeVisibility, label, value, onChange, id }: Props) {
    const showPassword = isVisible ? "text" : "password";
    const showIcon = isVisible ? <EyeOpenIcon className="h-4 w-4 " /> : <EyeClosedIcon className="h-4 w-4" />;

    return (
        <div className="flex flex-col space-y-1.5">
            <Label htmlFor="senha" className="text-white">{label}</Label>
            <div className="flex gap-3">
                <Input
                    id={id}
                    type={showPassword}
                    placeholder="Escreva a sua senha"
                    className="text-white"
                    value={value} // Adicione o valor aqui
                    onChange={onChange} // Conectar o onChange aqui
                />
                <Button onClick={onChangeVisibility} type="button" variant="ghost" size="icon" className="text-white bg-blue-900">
                    {showIcon}
                </Button>
            </div>
        </div>
    );
}
