import React, { useEffect, useState } from "react";
import { SafeAreaView, Text, StyleSheet, ActivityIndicator, FlatList, View, TouchableOpacity } from "react-native";
import { useRoute, useNavigation } from "@react-navigation/native";
import { SingleShoppingListRouteProp } from "../Navigation/NavigationTypes";
import { client, ShoppingList, FoodItem, DeleteShoppingListInput, DeleteItemInput } from "../Service/Api";

const SingleShoppingListScreen: React.FC = () => {
    const route = useRoute<SingleShoppingListRouteProp>();
    const navigation = useNavigation();
    const { input } = route.params;
    const [deleteList, setDeleteList] = useState<{ name: string }[]>([]);

    const [shoppingList, setShoppingList] = useState<ShoppingList | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchShoppingList = async () => {
            try {
                const data = await client.fetchShoppingListByName(input);
                setShoppingList(data);
            } catch (error) {
                console.error("Error fetching the shopping list", error);
            } finally {
                setLoading(false);
            }
        };

        fetchShoppingList();
    }, [input]);

    const handleDeletePress = async () => {
        try {
            const userId = await client.getUserIdHelper();
            const deleteRequest: DeleteShoppingListInput = {
                user_id: userId,
                from: input.from,
                to: input.to,
                name: shoppingList?.name,
            };
            await client.deleteShoppingList(deleteRequest);
            navigation.goBack();
        } catch (error) {
            console.error("Error deleting Shopping list", error);
        }
    };

    const handleItemDelete = (index: number) => {
        if (!shoppingList) return;
    
        const itemToDelete = shoppingList.items[index];
    
        setDeleteList((prevList) => [...prevList, { name: itemToDelete.name }]);
    
        const updatedItems = shoppingList.items.filter((_, i) => i !== index);
        setShoppingList({ ...shoppingList, items: updatedItems });
    };
    
    

    const handleSavePress = async () => {
        try {
            if (!shoppingList) return;

            const delete_items_request : DeleteItemInput = {
                from: input.from,
                to: input.to,
                name: shoppingList?.name,
                items: deleteList
            }

            await client.deleteItems(delete_items_request);
            alert("Shopping list saved successfully!");
        } catch (error) {
            console.error("Error saving the shopping list", error);
        }
    };

    const formatDate = (date: string) => {
        const newDate = new Date(date);
        return newDate.toLocaleDateString("fr-FR");
    };

    if (loading) {
        return (
            <SafeAreaView style={styles.loadingContainer}>
                <ActivityIndicator size="large" color="blue" />
            </SafeAreaView>
        );
    }

    if (!shoppingList) {
        return (
            <SafeAreaView style={styles.container}>
                <Text style={styles.errorText}>Failed to load the shopping list.</Text>
            </SafeAreaView>
        );
    }

    const formattedFrom = shoppingList.from
        ? new Date(Date.parse(shoppingList.from)).toLocaleDateString()
        : "Date unavailable";
    const formattedTo = shoppingList.to
        ? new Date(Date.parse(shoppingList.to)).toLocaleDateString()
        : "Date unavailable";

    const renderItem = ({ item, index }: { item: FoodItem; index: number }) => (
        <View style={styles.itemContainer}>
            <View>
                <Text style={styles.itemName}>{item.name}</Text>
                <Text style={styles.itemQuantity}>
                    {item.quantity} {item.unit}
                </Text>
            </View>
            <TouchableOpacity
                style={styles.deleteItemButton}
                onPress={() => handleItemDelete(index)}
            >
                <Text style={styles.deleteItemButtonText}>Supprimer</Text>
            </TouchableOpacity>
        </View>
    );

    return (
        <SafeAreaView style={styles.container}>
            <Text style={styles.title}>{shoppingList.name}</Text>

            <Text style={styles.dateRange}>
                From: {formattedFrom} To: {formattedTo}
            </Text>
            <FlatList
                data={shoppingList.items}
                renderItem={renderItem}
                keyExtractor={(item, index) => index.toString()}
                contentContainerStyle={styles.listContainer}
            />
            <View style={styles.buttonContainer}>
                <TouchableOpacity onPress={handleDeletePress} style={styles.deleteButton}>
                    <Text style={styles.buttonText}>Supprimer la liste</Text>
                </TouchableOpacity>
                <TouchableOpacity onPress={handleSavePress} style={styles.saveButton}>
                    <Text style={styles.buttonText}>Sauvegarder</Text>
                </TouchableOpacity>
            </View>
        </SafeAreaView>
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
        paddingHorizontal: 24,
        paddingVertical: 16,
        paddingBottom: 48,
        backgroundColor: 'white',
    },

    title: {
        fontSize: 24,
        fontWeight: "bold",
        marginBottom: 8,
        textAlign: "center",
    },
    dateRange: {
        fontSize: 16,
        color: '#666',
        marginBottom: 16,
        textAlign: "center",
    },
    listContainer: {
        paddingBottom: 80,
    },

    itemContainer: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        paddingVertical: 12,
        paddingHorizontal: 16,
        marginBottom: 8,
        backgroundColor: 'white', // Background for contrast
        borderRadius: 8,
        borderWidth: 1,
        borderColor: '#DDD',
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.2,
        shadowRadius: 1.41,
        elevation: 2,
    },
    itemName: {
        fontSize: 16,
        color: '#333',
    },
    itemQuantity: {
        fontSize: 16,
        color: '#666',
    },
    deleteItemButton: {
        backgroundColor: 'red',
        paddingVertical: 6,
        paddingHorizontal: 12,
        borderRadius: 4,
    },
    deleteItemButtonText: {
        color: 'white',
        fontWeight: 'bold',
        fontSize: 14,
    },
    errorText: {
        fontSize: 18,
        color: 'red',
        textAlign: 'center',
    },
    deleteButton: {
        backgroundColor: '#FFCC00',
        paddingVertical: 12,
        paddingHorizontal: 24,
        borderRadius: 5,
        alignItems: 'center',
        marginBottom: 24,
    },
    saveButton: {
        backgroundColor: 'green',
        paddingVertical: 12,
        paddingHorizontal: 24,
        borderRadius: 5,
        alignItems: 'center',
        marginBottom: 24,
    },
    buttonText: {
        color: 'black',
        fontWeight: 'bold',
    },
    buttonContainer: {
        paddingBottom: 24,
        paddingTop: 16,
        alignItems: 'center',
    },

});

export default SingleShoppingListScreen;
