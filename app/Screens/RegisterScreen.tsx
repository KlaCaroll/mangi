import React, { useState } from "react";
import { View, TextInput, Text, Pressable, StyleSheet } from "react-native";
import { client, RegisterInput } from "../Service/Api";
import { useNavigation } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import * as Yup from "yup";
import { RegisterValidationSchema } from "../Validation/Validation";

type RootStackParamList = {
  Auth: undefined;
  Login: undefined;
  Register: undefined;
  Home: undefined;
};

type AuthScreenNavigationProp = StackNavigationProp<RootStackParamList, 'Register'>;

export function RegisterScreen() {
  const navigation = useNavigation<AuthScreenNavigationProp>();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const handleRegister = async () => {
    const input: RegisterInput = { email, password, name };
       try {
         await RegisterValidationSchema.validate(input, { abortEarly: false });
         const res = await client.register(input);
         
         if (res.err) {
           if (res.err === "an user with this email already exists") {
              setError("Cet utilisateur existe déjà");
           } else { 
             if (res.err === "internal problem with database") {
               setError("Une erreur est survenue");
             }
             
             setError(res.err);
           }
           setSuccessMessage("");
         } else {
            navigation.navigate("Login");
           setSuccessMessage("Vous vous êtes inscrit avec succès !");
           setError("");
         }
       } catch (validationError) {
         if (validationError instanceof Yup.ValidationError) {
           setError(validationError.errors.join("\n"));
           setSuccessMessage("");
         } else {
           setError("Échec de l'inscription");
           setSuccessMessage("");
         }
       }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Inscription</Text>
      <TextInput
        style={styles.input}
        placeholder="Nom"
        value={name}
        onChangeText={setName}
        placeholderTextColor="#A9A9A9"
      />
      <TextInput
        style={styles.input}
        placeholder="Email"
        value={email}
        onChangeText={setEmail}
        placeholderTextColor="#A9A9A9"
      />
      <TextInput
        style={styles.input}
        placeholder="Mot de passe"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
        placeholderTextColor="#A9A9A9"
      />
      {error ? <Text style={styles.errorText}>{error}</Text> : null}
      {successMessage ? (
        <Text style={styles.successText}>{successMessage}</Text>
      ) : null}
      <Pressable style={styles.button} onPress={handleRegister}>
        <Text style={styles.buttonText}>Inscription</Text>
      </Pressable>

      <Pressable style={styles.backButton} onPress={() => navigation.navigate("Auth")}>
        <Text style={styles.backButtonText}>Retour</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#FFF5E5",
    padding: 20,
  },
  title: {
    fontSize: 28,
    fontWeight: "bold",
    color: "#000",
    marginBottom: 40,
  },
  input: {
    width: "90%",
    height: 50,
    borderColor: "#D3D3D3",
    borderWidth: 1,
    borderRadius: 10,
    paddingHorizontal: 15,
    backgroundColor: "#FFFFFF",
    fontSize: 16,
    color: "#000",
    marginBottom: 20,
  },
  button: {
    width: "90%",
    height: 50,
    backgroundColor: "#FFD700",
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 10,
    marginTop: 20,
  },
  buttonText: {
    fontSize: 18,
    fontWeight: "bold",
    color: "#000",
  },
  errorText: {
    color: "red",
    marginTop: 10,
  },
  successText: {
    color: "green",
    marginTop: 10,
  },
  backButton: {
    width: "90%",
    height: 50,
    borderColor: "#FFD700",
    borderWidth: 2,
    backgroundColor: "#FFFFFF",
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 10,
    marginTop: 20,
  },

  backButtonText: {
    fontSize: 18,
    fontWeight: "bold",
    color: "#000",
  },
});
