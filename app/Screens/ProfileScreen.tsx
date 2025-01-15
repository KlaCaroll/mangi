import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  StyleSheet,
  ActivityIndicator,
  ScrollView,
  TouchableOpacity,
} from "react-native";
import { client, ShowUserOutput } from "../Service/Api";
import Icon from "react-native-vector-icons/FontAwesome";
import { useNavigation, NavigationProp } from "@react-navigation/native";
import { ProfileToHomesStackParamList } from "../Navigation/NavigationTypes";
import LogoutButton from "../Components/Common/LogoutButton";

export function ProfileScreen() {
  const navigation = useNavigation<NavigationProp<ProfileToHomesStackParamList>>();
  const [user, setUser] = useState<ShowUserOutput | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [showUstensils, setShowUstensils] = useState(false);
  const [showPreferences, setShowPreferences] = useState(false);
  const [showPersonalInfo, setShowPersonalInfo] = useState(false);

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const userId = await client.getUserIdHelper();
        const data = await client.showProfile(userId);
        if (data.err === "internal problem with database") {
          setError("Une erreur est survenue");
          return;
        }
        setUser(data);
      } catch (error) {
        setError("Erreur lors de la récupération du profil");
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, []);

  if (loading) {
    return (
      <View style={styles.container}>
        <ActivityIndicator size="large" color="#FFD700" />
      </View>
    );
  }

  if (error) {
    return (
      <View style={styles.container}>
        <Text style={styles.errorText}>{error}</Text>
      </View>
    );
  }

  return (
    <ScrollView style={styles.container}>
      <View style={styles.profileHeader}>
        <Icon name="user-circle" size={50} color="#FFD700" />
        <Text style={styles.title}>{user?.name}</Text>
      </View>


      <TouchableOpacity onPress={() => setShowUstensils(!showUstensils)}>
        <Text style={styles.subtitle}>
          Ustensiles{" "}
          <Text style={styles.toggleSymbol}>{showUstensils ? "-" : "+"}</Text>
        </Text>
      </TouchableOpacity>
      {showUstensils &&
        (user?.ustensils && user.ustensils.length > 0 ? (
          user.ustensils.map((ustensil) => (
            <View key={ustensil.ustensil_id} style={styles.itemContainer}>
              <Text style={styles.item}>{ustensil.ustensil_name}</Text>
            </View>
          ))
        ) : (
          <View style={styles.itemContainer}>
            <Text style={styles.item}>Pas d'ustensiles</Text>
          </View>
        ))}

      <TouchableOpacity onPress={() => setShowPreferences(!showPreferences)}>
        <Text style={styles.subtitle}>
          Préférences{" "}
          <Text style={styles.toggleSymbol}>{showPreferences ? "-" : "+"}</Text>
        </Text>
      </TouchableOpacity>
      {showPreferences &&
        (user?.preferences && user.preferences.length > 0 ? (
          user.preferences.map((preference) => (
            <View key={preference.preference_id} style={styles.itemContainer}>
              <Text style={styles.item}> {preference.preference_name}</Text>
            </View>
          ))
        ) : (
          <View style={styles.itemContainer}>
            <Text style={styles.item}>Pas de préférences</Text>
          </View>
        ))}

      <TouchableOpacity onPress={() => setShowPersonalInfo(!showPersonalInfo)}>
        <Text style={styles.subtitle}>
          Informations personnelles{" "}
          <Text style={styles.toggleSymbol}>{showPersonalInfo ? "-" : "+"}</Text>
        </Text>
      </TouchableOpacity>
      {showPersonalInfo && (
        <View style={styles.personalInfoContainer}>
          <Text style={styles.infoItem}>Nom : {user?.name}</Text>
          <Text style={styles.infoItem}>Email : {user?.email}</Text>
        </View>
      )}

      <TouchableOpacity
        style={styles.homeButton}
        onPress={() => navigation.navigate("Homes", { userName: user?.name })}
      >
        <Text style={styles.buttonText}>Voir mes homes</Text>
      </TouchableOpacity>
      <LogoutButton />
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#FFF5E5",
    padding: 20,
  },
  profileHeader: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 20,
  },
  title: {
    fontSize: 28,
    fontWeight: "bold",
    marginLeft: 10,
    color: "#000",
  },
  label: {
    fontSize: 18,
    marginBottom: 10,
    color: "#000",
  },
  subtitle: {
    fontSize: 20,
    fontWeight: "bold",
    marginTop: 20,
    color: "#000",
  },
  itemContainer: {
    padding: 10,
    backgroundColor: "#f9f9f9",
    borderRadius: 8,
    marginBottom: 10,
    borderColor: "#ddd",
    borderWidth: 1,
  },
  item: {
    fontSize: 16,
    color: "#333",
  },
  toggleSymbol: {
    color: "#FFD700",
  },
  personalInfoContainer: {
    backgroundColor: "#f9f9f9",
    padding: 15,
    borderRadius: 8,
    marginTop: 10,
    borderWidth: 1,
    borderColor: "#ddd",
  },
  infoItem: {
    fontSize: 16,
    color: "#333",
    marginBottom: 10,
  },
  homeButton: {
    backgroundColor: "#FFD700",
    padding: 15,
    borderRadius: 10,
    marginTop: 20,
    alignItems: "center",
  },
  buttonText: {
    fontSize: 18,
    color: "#fff",
    fontWeight: "bold",
  },
  errorText: {
    color: "red",
    marginTop: 10,
  },
});
