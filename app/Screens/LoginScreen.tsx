import React, { useState } from "react";
import { View, Pressable, TextInput, Text, StyleSheet } from "react-native";
import { client, LoginInput } from "../Service/Api";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { useNavigation } from "@react-navigation/native";
import { StackNavigationProp } from "@react-navigation/stack";
import { LoginValidationSchema } from "../Validation/Validation";
import { ValidationError } from "yup";

type RootStackParamList = {
  Auth: undefined;
  Login: undefined;
  Register: undefined;
  Home: undefined;
  Main: undefined;
};

type AuthScreenNavigationProp = StackNavigationProp<RootStackParamList, "Login">;

export function LoginScreen() {
  const navigation = useNavigation<AuthScreenNavigationProp>();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const handleLogin = async () => {
    const input: LoginInput = { email, password };

    try {
      await LoginValidationSchema.validate(input, { abortEarly: false });
      const res = await client.login(input);

      if (res.err) {
        if (res.err === "no user found") {
          setError("Aucun utilisateur n'est associé à ce mail");
        } else if (res.err === "this password doesn't match") {
          setError("Mot de passe incorrect");
        } else if (res.err === "internal problem with database") {
          setError("Une erreur est survenue");
        }
        setSuccessMessage("");
      } else {
        await AsyncStorage.setItem("userToken", res.token);
        navigation.navigate("Main");
        setSuccessMessage("Vous vous êtes connecté avec succès !");
        setError("");
      }
    } catch (error) {
      if (error instanceof ValidationError) {
        setError("Validation error: " + error.errors.join(", "));
      } else {
        setError("Échec de la connexion");
      }
      setSuccessMessage("");
    }
  };

  return (
    <View style={styles.container}>
      <View
        accessible={true}
        accessibilityRole="header"
        focusable={true}
        accessibilityLabel="Titre : Connexion"
      >
        <Text style={styles.title}>Connexion</Text>
      </View>
      <TextInput
        style={styles.input}
        placeholder="Adresse e-mail (example@domaine.com)"
        value={email}
        onChangeText={setEmail}
        placeholderTextColor="#A9A9A9"
        accessible={true}
        accessibilityLabel="Champ email"
        keyboardType="email-address"
        textContentType="emailAddress"
        returnKeyType="next"
      />
      <TextInput
        style={styles.input}
        placeholder="Mot de Passe"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
        placeholderTextColor="#A9A9A9"
        accessible={true}
        accessibilityLabel="Champ mot de passe"
        textContentType="password"
        returnKeyType="done"
      />
      {error ? (
        <Text style={styles.errorText} accessibilityLiveRegion="polite">
          {error}
        </Text>
      ) : null}
      {successMessage ? (
        <Text style={styles.successText} accessibilityLiveRegion="polite">
          {successMessage}
        </Text>
      ) : null}
      <Pressable
        style={styles.button}
        onPress={handleLogin}
        accessible={true}
        accessibilityRole="button"
        accessibilityLabel="Bouton de connexion"
      >
        <Text style={styles.buttonText}>Connexion</Text>
      </Pressable>
      <Pressable
        style={styles.backButton}
        onPress={() => navigation.navigate("Auth")}
        accessible={true}
        accessibilityRole="button"
        accessibilityLabel="Bouton retour"
      >
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
    marginBottom: 40,
    color: "#000",
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
