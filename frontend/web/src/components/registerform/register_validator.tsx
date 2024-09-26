import * as yup from "yup";

export const formSchema = yup.object().shape({
  nome: yup.string().required("Nome é obrigatório"),
  email: yup.string().email("Email inválido").required("Email é obrigatório"),
  senha: yup.string().min(6, "A senha deve ter pelo menos 6 caracteres").required("Senha é obrigatória"),
  confirmaSenha: yup.string()
    .oneOf([yup.ref('senha')], 'As senhas devem corresponder')
    .required("Confirmação de senha é obrigatória"),
  telefone: yup.string()
    .required("Telefone é obrigatório")
    .matches(/^\(\d{2}\) \d{5}-\d{4}$/, "Telefone deve estar no formato (xx) xxxxx-xxxx"), // Regex para validar o telefone
});
