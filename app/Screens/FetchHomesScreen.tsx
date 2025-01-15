import React, { useEffect, useState } from "react";
import { client, Home } from "../Service/Api";
import {
  StyleSheet,
  Text,
  View,
  SafeAreaView,
  FlatList,
  ActivityIndicator,
  TouchableOpacity,
  Alert,
} from "react-native";
import { useNavigation, RouteProp, useRoute, NavigationProp } from "@react-navigation/native";
import { ProfileToHomesStackParamList } from "../Navigation/NavigationTypes";
import Icon from "react-native-vector-icons/FontAwesome";

type HomesRouteProp = RouteProp<ProfileToHomesStackParamList, "Homes">;
type HomesNavigationProp = NavigationProp<ProfileToHomesStackParamList>;

export function FetchHomesScreen() {
  const route = useRoute<HomesRouteProp>();
  const userName = route.params?.userName;
  const [homes, setHomes] = useState<Home[]>([]);
  const navigation = useNavigation<HomesNavigationProp>();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [refresh, setRefresh] = useState(false);

  const fetchHomes = async () => {
    try {
      const response = await client.fetchHomes();
      setHomes(response);      
    } catch (error) {
      setError("Erreur lors de la récupération des maisons");
    } finally {
      setLoading(false);
    }
  };

  const deleteHome = async (home_name: string) => {
    try {
      await client.deleteHome(home_name);
      setHomes(homes.filter((home) => home.name !== home_name));
    } catch (error) {
      Alert.alert("Erreur", "Erreur lors de la suppression du foyer");
    }
  };

  useEffect(() => {
    fetchHomes();
  }, [refresh]);

  const renderHomeItem = ({ item }: { item: Home }) => (
    <View style={styles.cardContainer}>
      <TouchableOpacity
        onPress={() =>
          navigation.navigate("SingleHome", { home_name: item.name, home_id: item.id })
        }
        style={{ flex: 1 }}
      >
        <View style={styles.card}>
          <View style={styles.cardContent}>
            <Text style={styles.homeName}>{item.name}</Text>
            <Text style={styles.homeOwner}>Propriétaire: {item.owner_name}</Text>
          </View>
          <TouchableOpacity
            onPress={() => deleteHome(item.name)}
            style={styles.deleteButton}
          >
            <Icon name="trash" size={24} color="red" />
          </TouchableOpacity>
        </View>
      </TouchableOpacity>
    </View>
  );

  const renderEmptyComponent = () => (
    <View style={styles.emptyContainer}>
      <Text style={styles.emptyText}>Vous n'avez pas encore de foyer !</Text>
    </View>
  );

  if (loading) {
    return (
      <SafeAreaView style={styles.container}>
        <ActivityIndicator size="large" color="#FFD700" />
      </SafeAreaView>
    );
  }

  if (error) {
    return (
      <SafeAreaView style={styles.container}>
        <Text style={styles.errorText}>{error}</Text>
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      <TouchableOpacity
        onPress={() => navigation.goBack()}
        style={styles.backButton}
      >
        <Text style={styles.backButtonText}>Retour</Text>
      </TouchableOpacity>
      <TouchableOpacity
        onPress={() => navigation.navigate("CreateHome", { setRefresh })}
        style={styles.addButton}
      >
        <Text style={styles.addButtonText}>+ Ajouter un foyer</Text>
      </TouchableOpacity>

      <Text style={styles.title}>Mes foyers</Text>

      <FlatList
        data={homes}
        keyExtractor={(item) => item.id.toString()}
        renderItem={renderHomeItem}
        contentContainerStyle={styles.listContainer}
        ListEmptyComponent={renderEmptyComponent}
      />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 10,
    backgroundColor: "#FFF5E5",
  },
  listContainer: {
    paddingBottom: 20,
  },
  backButton: {
    padding: 10,
    marginTop: 20,
    marginLeft: 10,
    backgroundColor: "#FFD700",
    borderRadius: 5,
    marginBottom: 20,
    width: 100,
    alignItems: "center",
  },
  backButtonText: {
    fontSize: 16,
    color: "#fff",
    fontWeight: "bold",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    color: "#333",
    marginTop: 20,
    marginBottom: 20,
    textAlign: "center",
  },
  cardContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
  },
  card: {
    flex: 1,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    padding: 20,
    marginVertical: 10,
    backgroundColor: "#FFFFFF",
    borderRadius: 15,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 6,
    elevation: 6,
    borderWidth: 1,
    borderColor: "#DDD",
  },
  cardContent: {
    flex: 1,
  },
  homeName: {
    fontSize: 20,
    fontWeight: "bold",
    color: "#333",
    marginBottom: 5,
  },
  homeOwner: {
    fontSize: 16,
    color: "#555",
  },
  deleteButton: {
    padding: 10,
  },
  errorText: {
    color: "red",
    textAlign: "center",
    marginTop: 20,
  },
  emptyContainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    padding: 20,
  },
  emptyText: {
    fontSize: 18,
    color: "#A9A9A9",
  },
  addButton: {
    padding: 10,
    backgroundColor: "#FFD700",
    borderRadius: 5,
    alignItems: "center",
    marginTop: 20,
    alignSelf: "center",
    width: 200,
  },
  addButtonText: {
    color: "#FFFFFF",
    fontSize: 16,
    fontWeight: "bold",
  },
});