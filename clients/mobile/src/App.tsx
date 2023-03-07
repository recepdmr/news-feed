import {NavigationContainer} from '@react-navigation/native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import * as React from 'react';
import {View, Text} from 'react-native';
import Home from './screens/home/Home';

const DetailsScreen: React.FC = () => {
  return (
    <View>
      <Text>Details Screen</Text>
    </View>
  );
};

const screens: [{name: string; component: React.FC<{}>; title: string}] = [
  {
    name: 'Home',
    component: Home,
    title: 'UGM Mobile Application',
  },
  {
    name: 'Details',
    component: DetailsScreen,
    title: 'Detail screen',
  },
];

const defaultScreen = screens[0];

const Stack = createNativeStackNavigator();

function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName={defaultScreen.name}>
        {screens.map(screen => (
          <Stack.Screen
            key={screen.name}
            name={screen.name}
            component={screen.component}
            options={{title: screen.title}}
          />
        ))}
      </Stack.Navigator>
    </NavigationContainer>
  );
}

export default App;
