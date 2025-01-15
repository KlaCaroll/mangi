// import React, { useEffect, useState } from "react";
// import {
//   View,
//   Text,
//   TextInput,
//   TouchableOpacity,
//   StyleSheet,
//   SafeAreaView,
//   ScrollView,
//   Alert,
//   FlatList,
// } from "react-native";
// import { useNavigation } from "@react-navigation/native";
// import { client } from "../Service/Api";
// import SearchBar from "../Components/Common/SearchBar";

// interface Ingredient {
//   id: number;
//   name: string;
// }

// export function AddRecipeScreen() {
//   const [name, setName] = useState<string>("");
//   const [instructions, setInstructions] = useState<string>("");
//   const [ingredients, setIngredients] = useState<Ingredient[]>([]);
//   const [availableIngredients, setAvailableIngredients] = useState<
//     Ingredient[]
//   >([]);
//   const [searchTerm, setSearchTerm] = useState<string>("");
//   const navigation = useNavigation();

//   useEffect(() => {
//     const fetchIngredients = async () => {
//       try {
//         const res = await client.fetchIngredients(); // Assurez-vous que cette fonction existe
//         setAvailableIngredients(res.ingredients); // Assurez-vous que la réponse a la bonne structure
//       } catch (error) {
//         console.error("Erreur lors de la récupération des ingrédients:", error);
//       }
//     };

//     fetchIngredients();
//   }, []);

//   const handleAddIngredient = (ingredient: Ingredient) => {
//     if (!ingredients.some((ing) => ing.id === ingredient.id)) {
//       setIngredients([...ingredients, ingredient]);
//     }
//   };

//   const handleAddRecipe = async () => {
//     if (!name || ingredients.length === 0 || !instructions) {
//       Alert.alert("Erreur", "Tous les champs sont obligatoires.");
//       return;
//     }

//     try {
//       await client.createRecipe({ name, ingredients, instructions });
//       Alert.alert("Succès", "Recette ajoutée avec succès !");
//       navigation.goBack(); // Retourne à la liste des recettes
//     } catch (error) {
//       console.error("Erreur lors de l'ajout de la recette:", error);
//       Alert.alert(
//         "Erreur",
//         "Une erreur est survenue lors de l'ajout de la recette."
//       );
//     }
//   };

//   const filteredIngredients = availableIngredients.filter((ingredient) =>
//     ingredient.name.toLowerCase().includes(searchTerm.toLowerCase())
//   );

//   return (
//     <SafeAreaView style={styles.container}>
//       <ScrollView contentContainerStyle={styles.scrollContainer}>
//         <Text style={styles.label}>Nom de la recette</Text>
//         <TextInput
//           style={styles.input}
//           placeholder="Ex: Tarte aux pommes"
//           value={name}
//           onChangeText={setName}
//         />

//         <Text style={styles.label}>Ingrédients</Text>
//         <SearchBar
//           value={searchTerm}
//           onChangeText={setSearchTerm}
//           placeholder="Rechercher un ingrédient"
//         />
//         <FlatList
//           data={filteredIngredients}
//           keyExtractor={(item) => item.id.toString()}
//           renderItem={({ item }) => (
//             <TouchableOpacity onPress={() => handleAddIngredient(item)}>
//               <Text style={styles.ingredientItem}>{item.name}</Text>
//             </TouchableOpacity>
//           )}
//         />

//         <Text style={styles.label}>Instructions</Text>
//         <TextInput
//           style={[styles.input, styles.multiLineInput]}
//           placeholder="Ex: Mélanger les ingrédients, cuire au four..."
//           value={instructions}
//           onChangeText={setInstructions}
//           multiline
//         />

//         <TouchableOpacity style={styles.addButton} onPress={handleAddRecipe}>
//           <Text style={styles.addButtonText}>Ajouter la recette</Text>
//         </TouchableOpacity>
//       </ScrollView>
//     </SafeAreaView>
//   );
// }

// const styles = StyleSheet.create({
//   container: {
//     flex: 1,
//     backgroundColor: "#FFF5E5",
//   },
//   scrollContainer: {
//     padding: 20,
//   },
//   label: {
//     fontSize: 16,
//     fontWeight: "bold",
//     marginBottom: 5,
//     color: "#333",
//   },
//   input: {
//     backgroundColor: "#FFFFFF",
//     padding: 10,
//     borderRadius: 5,
//     borderWidth: 1,
//     borderColor: "#D3D3D3",
//     marginBottom: 15,
//   },
//   multiLineInput: {
//     height: 100,
//     textAlignVertical: "top",
//   },
//   addButton: {
//     padding: 15,
//     backgroundColor: "#FFA500",
//     borderRadius: 5,
//     alignItems: "center",
//     marginTop: 10,
//   },
//   addButtonText: {
//     color: "#FFFFFF",
//     fontSize: 16,
//     fontWeight: "bold",
//   },
//   ingredientItem: {
//     padding: 10,
//     borderBottomWidth: 1,
//     borderBottomColor: "#D3D3D3",
//   },
// });
