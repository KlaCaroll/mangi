import React, { useEffect, useState } from "react";
import {
  View,
  ScrollView,
  FlatList,
  StyleSheet,
  SafeAreaView,
  Text,
  TouchableOpacity,
} from "react-native";
import { useNavigation } from "@react-navigation/native";
import { client, Recipe } from "../Service/Api";
import { RecipesStackNavigationProp } from "../Navigation/NavigationTypes";
import RecipeCard from "../Components/Specific/RecipeCard";
import SearchBar from "../Components/Common/SearchBar";

export function RecipesScreen() {
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [suggestions, setSuggestions] = useState<Recipe[]>([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [selectedSuggestion, setSelectedSuggestion] = useState<string | null>(
    null
  );
  const navigation = useNavigation<RecipesStackNavigationProp>();

  const fetchRecipes = async (name: string) => {
    try {
      const res = await client.fetchRecipes({ name, preference: false });
      setRecipes(res.recipes);
    } catch (err) {
      console.error("Error fetching recipes:", err);
    }
  };

  useEffect(() => {
    fetchRecipes("");
  }, []);

  useEffect(() => {
    if (searchTerm) {
      fetchRecipes(searchTerm);
    } else {
      fetchRecipes("");
    }
  }, [searchTerm]);

  useEffect(() => {
    if (searchTerm) {
      const filtered = recipes?.filter((recipe) =>
        recipe.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSuggestions(filtered);
      setShowSuggestions(filtered.length > 0);
    } else {
      setSuggestions([]);
      setShowSuggestions(false);
    }
  }, [searchTerm, recipes]);

  const handleSuggestionPress = (name: string) => {
    setSearchTerm(name);
    setShowSuggestions(false);
    setSelectedSuggestion(name);
  };

  const filteredRecipes = recipes?.filter((recipe) =>
    recipe.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <ScrollView>
        <View style={styles.addButtonContainer}>
          <TouchableOpacity
            style={styles.addButton}
            //onPress={() => navigation.navigate("CreateRecipe")}
          >
            <Text style={styles.addButtonText}>+ Ajouter une recette</Text>
          </TouchableOpacity>
        </View>
        <View style={styles.container}>
          <SearchBar
            value={searchTerm}
            onChangeText={setSearchTerm}
            placeholder="Rechercher une recette"
            style={styles.searchBar}
          />
          {showSuggestions && suggestions.length > 0 && (
            <View style={styles.suggestionsContainer}>
              <FlatList
                data={suggestions}
                keyExtractor={(item) => item.id.toString()}
                renderItem={({ item }) => (
                  <TouchableOpacity
                    onPress={() => handleSuggestionPress(item.name)}
                    style={[
                      styles.suggestionItem,
                      selectedSuggestion === item.name &&
                        styles.selectedSuggestion,
                    ]}
                  >
                    <Text>{item.name}</Text>
                  </TouchableOpacity>
                )}
                style={styles.suggestionsList}
              />
            </View>
          )}
          <FlatList
            data={filteredRecipes}
            keyExtractor={(item) => item.id.toString()}
            renderItem={({ item }) => (
              <TouchableOpacity
                onPress={() =>
                  navigation.navigate("SingleRecipe", { id: item.id })
                }
                style={styles.recipeItem}
              >
                <RecipeCard
                  title={item.name}
                  description={item.description}
                />
              </TouchableOpacity>
            )}
            contentContainerStyle={styles.listContainer}
            initialNumToRender={10}
            windowSize={5}
          />
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  scrollViewContent: {
    flexGrow: 1,
    alignItems: "center",
    padding: 10,
  },
  container: {
    flex: 1,
    padding: 10,
    backgroundColor: "#FFF5E5",
    width: "100%",
    alignItems: "center",

  },

  searchBar: {
    marginBottom: 10,
    width: "100%",
  },
  suggestionsContainer: {
    position: "absolute",
    top: 60,
    left: 10,
    right: 10,
    backgroundColor: "#FFFFFF",
    borderColor: "#D3D3D3",
    borderWidth: 1,
    borderRadius: 5,
    zIndex: 1000,
  },
  suggestionsList: {
    maxHeight: 150,
  },
  suggestionItem: {
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: "#D3D3D3",
  },
  selectedSuggestion: {
    backgroundColor: "#FFD700",
  },
  listContainer: {
    paddingBottom: 20,
    width: "100%",
    alignItems: "center",
  },
  recipeItem: {

  },
  addButtonContainer: {
    backgroundColor: "#FFF5E5",
    width: "100%",
    alignItems: "center",
  },
  addButton: {
    padding: 15,
    backgroundColor: "#FFD700",
    borderRadius: 5,
    alignItems: "center",
    margin: 10,
    alignSelf: "center",
  },
  addButtonText: {
    color: "#FFFFFF",
    fontSize: 16,
    fontWeight: "bold",

  },
});
