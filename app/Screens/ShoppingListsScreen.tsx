import React, { useEffect, useState, useCallback } from "react";
import { client, ShoppingList } from "../Service/Api";
import { Alert, SafeAreaView, ActivityIndicator, StyleSheet, FlatList, Dimensions, Text, View, ScrollView } from "react-native";
import ListCard from "../Components/Specific/ListCard";
import { useNavigation, useFocusEffect } from "@react-navigation/native";
import { ShoppingListStackNavigationProp } from "../Navigation/NavigationTypes";
import SearchBar from "../Components/Common/SearchBar";

const ShoppingListsScreen: React.FC = () => {
    const [shoppingLists, setShoppingLists] = useState<ShoppingList[]>([]);
    const [loading, setLoading] = useState(true);
    const navigation = useNavigation<ShoppingListStackNavigationProp>();

    const screenWidth = Dimensions.get('window').width;
    const cardWidth = 180;
    const numColumns = Math.floor(screenWidth / cardWidth);

    const fetchShoppingLists = async () => {
        try {
            setLoading(true);
            const shoppingListsData = await client.fetchShoppingLists();
            setShoppingLists(shoppingListsData);
        } catch (error) {
            console.error("Error fetching Shopping lists", error);
            Alert.alert("Error", "Failed to recover shopping lists");
        } finally {
            setLoading(false);
        }
    };
    // this allow to re-fetch the list everytime we access the screen to reflect the deletion of a shopping list
    useFocusEffect(
        useCallback(() => {
            fetchShoppingLists();
        }, [])
    );

    const handleCardPress = (shoppingList: ShoppingList) => {
        const input = { from: shoppingList.from, to: shoppingList.to, name: shoppingList.name};
        navigation.navigate('SingleShoppingListScreen', { input: input });
    };

    if (loading) {
        return (
            <SafeAreaView style={styles.loadingContainer}>
                <ActivityIndicator size="large" color="blue" />
            </SafeAreaView>
        );
    }

    if (!shoppingLists){
        return (
            <View style={styles.emptyContainer}>
                <Text  style={styles.emptyText}> No Shopping lists available.</Text>
            </View>
        )
    }

    return (
        <View style={styles.container}>
        <ScrollView contentContainerStyle={styles.scrollContainer}>
          <Text style={styles.header}>Mes listes de courses</Text>
          <View style={styles.spacing} />
          <SearchBar />
          {shoppingLists.length === 0 ? (
            <View style={styles.emptyContainer}>
              <Text style={styles.emptyText}>No Shopping lists available.</Text>
            </View>
          ) : (
            <FlatList
              data={shoppingLists}
              renderItem={({ item, index }) => (
                <ListCard
                  shoppingList={{ ...item, name: `Liste ${index + 1}` }}
                  onPress={() => handleCardPress(item)}
                />
              )}
              keyExtractor={(item, index) => index.toString()}
              contentContainerStyle={styles.contentContainer}
            />
          )}
        </ScrollView>
      </View>
    );
  };


const styles = StyleSheet.create({
    loadingContainer: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",

    },
    container: {
        flex: 1,
        backgroundColor: 'white',
        paddingBottom : 60,

    },
    list: {
        flex: 1,
    },

    emptyContainer: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    emptyText: {
        fontSize: 18,
        color: "#FF7043",
    },

    header: {
        fontSize: 28,
        fontWeight: "bold",
        textAlign: "left",
        marginTop: 140,
        marginBottom: 20,
        color: "Black",
        marginLeft: 15,

    },
    spacing :{
        height : 80,

    },
    contentContainer: {
        paddingBottom: 100,
    },

    scrollContainer: {
        flexGrow: 1,
        paddingBottom: 100,
      },


});

export default ShoppingListsScreen;
