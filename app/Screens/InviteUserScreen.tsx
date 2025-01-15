import React, { useState } from "react";
import {
  View,
  Text,
  TextInput,
  Alert,
  StyleSheet,
  TouchableOpacity,
} from "react-native";
import { useRoute, RouteProp, useNavigation } from "@react-navigation/native";
import { client, HomeInvitationInput } from "../Service/Api";
import { ProfileToHomesStackParamList } from "../Navigation/NavigationTypes";
import { StackNavigationProp } from "@react-navigation/stack";


type InviteScreenRouteProp = RouteProp<ProfileToHomesStackParamList, "Invite">;
type InviteNavigationProp = StackNavigationProp<
  ProfileToHomesStackParamList,
  "SingleHome"
>;


export function InviteUserScreen() {
  const route = useRoute<InviteScreenRouteProp>();
  const { home_id } = route.params;
    const navigation = useNavigation<InviteNavigationProp>();
  const [invitation_to, setInvitationTo] = useState("");

  const handleInviteUser = async () => {
    if (home_id === undefined) {
      Alert.alert("Erreur", "L'ID de la maison est manquant");
      return;
    }
    const input: HomeInvitationInput = { home_id, invitation_to };
    try {
      const response = await client.inviteUserHome(input);
      if (response.err) {
        Alert.alert("Erreur", response.err);
      } else {
        const {name :home_name} = response;
        Alert.alert(
          "Invitation envoyée",
          `Invitation envoyée à ${invitation_to}`
        );
          navigation.navigate("SingleHome", { home_id, home_name, refresh: true });

      }
    } catch (error) {
      Alert.alert("Erreur", "Échec de l'envoi de l'invitation");
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Inviter une personne au foyer</Text>
      <TextInput
        placeholder="Email de l'invité"
        value={invitation_to}
        onChangeText={setInvitationTo}
        style={styles.input}
      />
      <TouchableOpacity style={styles.button} onPress={handleInviteUser}>
        <Text style={styles.buttonText}>Envoyer l'invitation</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    padding: 16,
    backgroundColor: "#FFF5E5",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 16,
    color: "#000",
  },
  input: {
    borderWidth: 1,
    padding: 10,
    marginBottom: 16,
    width: "100%",
    borderColor: "#FFD700",
    borderRadius: 10,
    backgroundColor: "#FFF",
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.2,
    shadowRadius: 5,
    elevation: 3,
    color: "#888",
  },
  button: {
    backgroundColor: "#FFD700",
    padding: 10,
    borderRadius: 5,
    alignItems: "center",
  },
  buttonText: {
    fontSize: 18,
    color: "#FFFF",
  },
});
