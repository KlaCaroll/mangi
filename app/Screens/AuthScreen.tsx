import React from "react";
import { View, Text, Pressable, StyleSheet, Image } from "react-native";
import { StackNavigationProp } from "@react-navigation/stack";
import { useNavigation, } from "@react-navigation/native";

type RootStackParamList = {
  Auth: undefined;
  Login: undefined;
  Register: undefined;
};


type AuthScreenNavigationProp = StackNavigationProp<RootStackParamList, 'Auth'>;


export function AuthScreen() {
  const navigation = useNavigation<AuthScreenNavigationProp>();

  return (

    <View style={styles.container}>
          <View style={styles.logoContainer}>
        <Image
          source={require("../Assets/Icon/logo_mangi.png")}
          style={styles.logo}
        />
      </View>

      <Text style={styles.title}>Bienvenue !</Text>
      <Pressable
        style={styles.button}
        onPress={() => navigation.navigate("Login")}
      >
        <Text style={styles.buttonText}>Connexion</Text>
      </Pressable>
      <Pressable
        style={styles.button}
        onPress={() => navigation.navigate("Register")}
      >
        <Text style={styles.buttonText}>Inscription</Text>
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
  logo: {

    resizeMode: 'contain',
  },


  button: {
    width: "80%",
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

  logoContainer:{
    position: "absolute",
    top: 200,
    alignItems: "center",
    width: "100%",


  },
});
