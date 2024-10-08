import {
    AlertDialog,
    AlertDialogContent,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogAction
  } from "@/components/ui/alert-dialog";
  import { useRouter } from "next/navigation";
  
  interface AlertLoginErrorsProps {
    show: boolean;
    onClose(open: boolean): void;
  }
  
  export function AlertLoginErrors({ show, onClose }: AlertLoginErrorsProps) {
    const router = useRouter();
    return (
      <div className="h-5/6 w-5/6 items-center justify-center bg-fixed">
        <AlertDialog open={show} onOpenChange={onClose}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Houve um erro no login, tente novamente</AlertDialogTitle>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogAction onClick={() => { router.push("/login") }} >Tentar Novamente</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
  
    );
  }
  