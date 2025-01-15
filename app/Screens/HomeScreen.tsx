import React from 'react';
import { StyleSheet, Text, View, SafeAreaView, ScrollView, Image } from 'react-native';
import Carousel from '../Components/Specific/Carrousel';
import SearchBar from '../Components/Common/SearchBar';

export function HomeScreen() {
  const recipes = [
    {
      image: require('../Assets/Images/pate_pesto.jpg'),
      title: 'Pâtes au pesto',
      description: 'Délicieuses farfales au pesto fait maison',
    },
    {
      image: require('../Assets/Images/caesar_salad.jpg'),
      title: 'Salade Caésar',
      description: 'Salade Caésar croquante avec sauce maison',
    },
    {
      image: require('../Assets/Images/pizza_margheritta.jpg'),
      title: 'Pizza Margherita',
      description: 'Pizza classique avec mozzarella fraîche',
    },
    {
      image: require('../Assets/Images/tarte_citron.jpg'),
      title: 'Tarte au citron',
      description: 'Tarte au citron meringuée maison',
    },
  ];

  const shoppingListPictures = [
    {
      image: require('../Assets/Images/shopping_list.jpg'),
      title: 'Repas de Noël ',
      description: 'Partagé avec : Lila, Nathan, Philippe',
    },
    {
      image: require('../Assets/Images/shopping_list.jpg'),
      title: 'Pot de départ Vivian',
      description: 'Partagé avec : Thomas, Benjamin, Maryam',
    },
    {
      image: require('../Assets/Images/shopping_list.jpg'),
      title: 'Pizza Margherita',
      description: 'Pizza classique avec mozzarella fraîche',
    },
    {
      image: require('../Assets/Images/shopping_list.jpg'),
      title: 'Tarte au citron',
      description: 'Tarte au citron meringuée maison',
    },
  ];




  return (
    <SafeAreaView style={styles.container}>
      <ScrollView>
        <View style={styles.avatarContainer}>
          <Image
            source={require('../Assets/Avatars/avatar_1.png')}
            style={styles.avatar}
          />
        </View>

        <View style={styles.searchBarContainer}>

          <SearchBar></SearchBar>
        </View>


        <View style={styles.carouselContainer}>

        <Text style={[styles.customFontText, { marginBottom: 30 }]}>Mes Recettes</Text>
          <Carousel recipes={recipes} />
        </View>

        <View style={styles.carouselContainer}>
          <Text style={[styles.customFontText, { marginBottom: 30 }]}>Mes Listes</Text>
          <Carousel recipes={shoppingListPictures} />
        </View>


      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'flex-start',
    alignItems: 'center',
    paddingTop: 50,
    backgroundColor: '#fff',
  },

  avatarContainer: {
    position: 'absolute',
    top: 40,
    right: 30,
    width: 75,
    height: 70,
    borderRadius: 30,
    backgroundColor: '#ddd',
    justifyContent: 'center',
    alignItems: 'center',
    overflow: 'hidden',
  },

  avatar: {
    width: '170%',
    height: '100%',
    resizeMode: 'cover',
  },

  carouselContainer: {
    marginTop: 50,
    width: '100%',


  },

  customFontText: {
    fontFamily: 'Happy Monkey',
    fontSize: 24,
  },

  searchBarContainer:{

    marginTop: 230,
    paddingHorizontal: 20,

  },
});
