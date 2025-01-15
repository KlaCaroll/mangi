import React from 'react';
import { View, Text, Image, StyleSheet, ImageSourcePropType } from 'react-native';

type RecipeCardProps = {
  source?: ImageSourcePropType;
  title: string;
  description: string;
  width? : number,
  titleColor?:string;
  descriptionColor?:string;

};

const RecipeCard: React.FC<RecipeCardProps> = ({
  source,
  title,
  description,
  width =266
}) => {
  return (
    <View style={styles.cardContainer}>
      <Image source={source} style={styles.image} />
      <View style={styles.infoContainer}>
        <Text style={styles.title}>{title}</Text>
        <Text style={styles.description}>{description}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  cardContainer: {
    width: 266,
    height: 229,
    borderWidth: 0.2,
    borderBlockColor: 'black',
    borderRadius: 10,
    backgroundColor: '#fff',
    marginTop: 10,
    marginRight : 5,

  },
  image: {
    width: '100%',
    height: 150,
    resizeMode: 'cover',
  },
  infoContainer: {
    padding: 10,
    height: 90,
    justifyContent: 'center',

  },
  title: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#333',
  },
  description: {
    fontSize: 16,
    color: 'black',
    marginTop: 4,
  },
});

export default RecipeCard;
