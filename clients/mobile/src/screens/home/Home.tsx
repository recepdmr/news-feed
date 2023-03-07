import React from 'react';
import {Button, Text, View} from 'react-native';
import BaseLayout from '../../layout/BaseLayout';
import Stories from './components/Stories';

const Home: React.FC = ({navigation}: any) => {
  return (
    <BaseLayout>
      <Stories />
      <View>
        <Text>Home Screen</Text>
        <Button
          title="Go to Details"
          onPress={() => navigation.navigate('Details')}
        />
      </View>
    </BaseLayout>
  );
};

export default Home;
