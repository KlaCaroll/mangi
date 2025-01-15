import * as Yup from "yup";

export const LoginValidationSchema = Yup.object().shape({
  email: Yup.string()
    .email("Veuillez entrer une adresse e-mail valide")
    .required("L'adresse e-mail est obligatoire"),
  password: Yup.string()
    .min(4, "Le mot de passe doit contenir au moins 4 caractères")
    .required("Le mot de passe est obligatoire"),
});

export const RegisterValidationSchema = Yup.object().shape({
  name: Yup.string()
    .min(2, "Le nom doit contenir au moins 2 caractères")
    .required("Le nom est obligatoire"),
  email: Yup.string()
    .email("Veuillez entrer une adresse e-mail valide")
    .required("L'adresse e-mail est obligatoire"),
  password: Yup.string()
    .min(4, "Le mot de passe doit contenir au moins 4 caractères")
    .required("Le mot de passe est obligatoire"),
});
