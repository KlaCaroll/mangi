import React from 'react';
import { View, Image, StyleSheet, Dimensions, ScrollView } from 'react-native';
import RecipeCard from './RecipeCard';

interface CarouselProps {
   recipes :{image : any; title : string; description: string}[];
}


const Carousel: React.FC<CarouselProps> = ({ recipes }) => {
    return (
        <View style={styles.container}>
            <ScrollView
                horizontal
                showsHorizontalScrollIndicator={false}
                pagingEnabled
                decelerationRate="fast"
                style={styles.scrollView}
                contentContainerStyle={styles.scrollContainer}
            >

                {recipes.map((recipe, index) => (
                    <View key={index} style={styles.cardWrapper}>
                    <RecipeCard
                    key ={index}
                    source={recipe.image}
                    title={recipe.title}
                    description={recipe.description}
                    />
                    </View>

                ))}

            </ScrollView>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        height: 220,
        justifyContent: 'center',
    },
    image: {
        width: 150,
        height: 150,
        resizeMode: 'cover',
        marginHorizontal: 10,
        borderRadius: 10,
    },
    scrollView: {
        flexDirection: 'row',
    },
    scrollContainer: {
        alignItems: 'center',
    },

    cardWrapper:{
        marginRight :20,
    }
});

export default Carousel;
