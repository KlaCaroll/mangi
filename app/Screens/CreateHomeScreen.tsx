import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  TouchableOpacity,
  StyleSheet,
  SafeAreaView,
  Alert,
} from "react-native";
import { useNavigation, useRoute, RouteProp } from "@react-navigation/native";
import { client } from "../Service/Api";
import { ProfileToHomesStackParamList } from "../Navigation/NavigationTypes";

type CreateHomeScreenRouteProp = RouteProp<
  ProfileToHomesStackParamList,
  "CreateHome"
>;

export function CreateHomeScreen() {
  const [homeName, setHomeName] = useState<string>("");
  const navigation = useNavigation();
  const route = useRoute<CreateHomeScreenRouteProp>();
    

  const createHome = async () => {
    if (!homeName.trim()) {
      Alert.alert("Erreur", "Le nom du foyer ne peut pas être vide");
      return;
    }
    try {
      await client.createHome(homeName);
      Alert.alert("Succès", "Foyer créé avec succès");
      if (route.params?.setRefresh) {
        route.params.setRefresh((prev: boolean) => !prev);
      }
      navigation.goBack();
    } catch (error) {
      Alert.alert("Erreur", "Erreur lors de la création du foyer");
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <Text style={styles.title}>Créer un nouveau foyer</Text>
      <TextInput
        style={styles.input}
        placeholder="Nom du foyer"
        value={homeName}
        onChangeText={setHomeName}
      />
      <TouchableOpacity onPress={createHome} style={styles.button}>
        <Text style={styles.buttonText}>Créer</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#FFF5E5",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    color: "#333",
    marginBottom: 20,
    textAlign: "center",
  },
  input: {
    height: 40,
    borderColor: "#D3D3D3",
    borderWidth: 1,
    borderRadius: 5,
    paddingHorizontal: 10,
    marginBottom: 20,
    backgroundColor: "#F0F0F0",
    width: "80%",
    alignSelf: "center",
    // borderColor: "#D3D3D3", // Couleur de la bordure
    // borderWidth: 1, // Largeur de la bordure
  },
  button: {
    backgroundColor: "#FFD700",
    padding: 15,
    borderRadius: 5,
    alignItems: "center",
    alignSelf: "center",
    width: "50%",
  },
  buttonText: {
    color: "#FFFFFF",
    fontSize: 16,
    fontWeight: "bold",
  },
});
