import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, ActivityIndicator, FlatList,TouchableOpacity} from "react-native";
import { RouteProp, useRoute , useNavigation} from "@react-navigation/native";
import { ProfileToHomesStackParamList } from "../Navigation/NavigationTypes";
import { StackNavigationProp } from "@react-navigation/stack"; 
import { client, SingleHome } from "../Service/Api";
import Icon from "react-native-vector-icons/FontAwesome";

type SingleHomeRouteProp = RouteProp<
  ProfileToHomesStackParamList,
  "SingleHome"
  >;
type InviteNavigationProp = StackNavigationProp<
  ProfileToHomesStackParamList,
  "Invite"
>;



export function SingleHomeScreen() {
  const route = useRoute<SingleHomeRouteProp>();
  const navigation = useNavigation<InviteNavigationProp>();
  const { home_name } = route.params;
  const [home, setHome] = useState<SingleHome | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchHome = async () => {
      try {
        const response = await client.fetchSingleHome(home_name);
        if (response.err) {
          setError(response.err);
          return;
        }
        setHome(response);
      } catch (error) {
        setError("Erreur lors de la récupération de la maison");
      } finally {
        setLoading(false);
      }
    };
    fetchHome();
  }, [home_name]);

  


  

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
    <View style={styles.container}>
      <View style={styles.homeInfoContainer}>
        <Text style={styles.title}>Nom du foyer : {home?.name}</Text>
        <Text style={styles.description}>Propriétaire: {home?.owner_name}</Text>
      </View>

      <View style={styles.membersContainer}>
        <View style={styles.membersHeader}>
          <Text style={styles.subtitle}>Membres:</Text>
          <TouchableOpacity
            style={styles.button}
            onPress={() => navigation.navigate("Invite", { home_id: home?.id })}
          >
            <Text style={styles.buttonText}>Ajouter</Text>
          </TouchableOpacity>
        </View>

        <FlatList
          data={home?.members ?? []}
          keyExtractor={(item) => item.id.toString()}
          renderItem={({ item }) => (
            <View style={styles.card}>
              <Icon name="user" size={30} color="#FFD700" />
              <Text style={styles.member}>{item.name}</Text>
            </View>
          )}
          ListEmptyComponent={
            <Text style={styles.member}>Aucun membre trouvé</Text>
          }
        />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#FFF5E5",
  },
  homeInfoContainer: {
    marginTop: 100,
  },
  title: {
    fontSize: 28,
    fontWeight: "bold",
    marginBottom: 10,
    color: "#000",
  },
  description: {
    fontSize: 18,
    color: "#555",
    marginBottom: 20,
  },
  subtitle: {
    fontSize: 20,
    fontWeight: "bold",
    marginTop: 20,
    color: "#000",
  },
  membersContainer: {
    marginTop: 30,
  },
  card: {
    flexDirection: "row",
    alignItems: "center",
    backgroundColor: "#FFF",
    padding: 15,
    marginVertical: 10,
    borderRadius: 10,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.2,
    shadowRadius: 5,
    elevation: 3,
  },
  member: {
    fontSize: 16,
    marginLeft: 10,
    color: "#000",
  },
  errorText: {
    color: "red",
    textAlign: "center",
    marginTop: 20,
  },
  membersHeader: {
    flexDirection: "row",
    alignItems: "center",
    marginBottom: 8,
    justifyContent: "space-between",
  },
  button: {
    backgroundColor: "#FFD700",
    padding: 10,
    borderRadius: 5,
    marginLeft: 10,
  },
  buttonText: {
    fontSize: 16,
    color: "#FFF",
  },
});

