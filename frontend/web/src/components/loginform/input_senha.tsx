
import { EyeOpenIcon, EyeClosedIcon} from "@radix-ui/react-icons";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

interface Props {
    isVisible: boolean;
    onChangeVisibility: () => void;
}
export function Senhainput({ isVisible, onChangeVisibility }: Props) {
    const showPassword =  isVisible ? "text" : "password";
    const showIcon = isVisible ? <EyeOpenIcon className="h-4 w-4 " /> : <EyeClosedIcon className="h-4 w-4" />;
    return (
        <div className="flex flex-col space-y-1.5" >
            <Label htmlFor="senha" className="text-white">Senha</Label>
            <div className="flex gap-3">
                <Input id="senha" type={showPassword} placeholder="Escreva o sua senha" className="text-white" />
                <Button variant="ghost" size="icon" onClick={onChangeVisibility} type="button" className="text-white bg-blue-900">
                    {showIcon}
                </Button>
            </div>
        </div>
    )
}