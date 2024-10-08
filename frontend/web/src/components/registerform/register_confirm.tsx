import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogAction
} from "@/components/ui/alert-dialog";
import { useRouter } from "next/navigation";

interface AlertRegisterConfirmProps {
  show: boolean;
  onClose(open: boolean): void;
}

export function AlertRegisterConfirm({ show, onClose }: AlertRegisterConfirmProps) {
  const router = useRouter();
  return (
    <div className="h-5/6 w-5/6 items-center justify-center bg-fixed">
      <AlertDialog open={show} onOpenChange={onClose}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Cadastro Feito com Sucesso</AlertDialogTitle>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogAction onClick={() => { router.push("/login") }} >Logar</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>

  );
}
