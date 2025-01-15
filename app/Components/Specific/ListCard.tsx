import React from 'react';
import { View, Text, FlatList, StyleSheet, TouchableOpacity } from 'react-native';
import { ShoppingList } from '../../Service/Api';

interface ListCardProps {
    shoppingList: ShoppingList;
    onPress: () => void;
}

const ListCard: React.FC<ListCardProps> = ({ shoppingList,onPress }) => {
    const displayItems = shoppingList.items.slice(0, 3);
    const hasMoreItems = shoppingList.items.length > 3;

    const formatDate = (date: string) => {
        const newDate = new Date(date);
        return newDate.toLocaleDateString("fr-FR");
    };

    return (
        <TouchableOpacity style={styles.card} onPress={onPress}>
            <View style={styles.textContainer}>
            <Text style={styles.listTitle}>{shoppingList.name}</Text>
            <Text style={styles.listSubtitle}>
                Du: {formatDate(shoppingList.from)} Au: {formatDate(shoppingList.to)}
                </Text>
            </View>
            <View style={styles.itemsContainer}>
                {displayItems.map((foodItem, index) => (
                    <Text key={index} style={styles.foodItemText}>

                    </Text>
                ))}
                {hasMoreItems && <Text style={styles.moreText}>...</Text>}
            </View>
        </TouchableOpacity>
    );
};

const styles = StyleSheet.create({
    card: {
        backgroundColor: '#ffffff',
        borderRadius: 12,
        padding: 10,
        marginVertical: 10,
        shadowColor: '#000',
        shadowOpacity: 0.2,
        shadowOffset: { width: 0, height: 2 },
        shadowRadius: 6,
        elevation: 3,
        width: 360,
        height : 80,
        flexDirection: 'column',
        alignSelf : 'center',
    },
    listTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        color: '#333',
        marginBottom: 4,
    },
    listSubtitle: {
        fontSize: 12,
        color: '#555',

    },
    itemsContainer: {
        flexDirection: 'row',
        alignItems : 'center',
    },
    foodItemText: {
        fontSize: 14,
        color: '#444',
        marginBottom: 2,
    },
    moreText: {
        fontSize: 12,
        color: '#888',
        fontStyle: 'italic',
    },

    textContainer:{
        flexDirection:'column',
        marginBottom: 8,

    },
});

export default ListCard;
