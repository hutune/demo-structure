import 'package:flutter/material.dart';

void main() {
  runApp(const RmnApp());
}

class RmnApp extends StatelessWidget {
  const RmnApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'RMN Platform',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
        useMaterial3: true,
      ),
      home: const Scaffold(
        body: Center(
          child: Text('RMN Platform - Coming Soon'),
        ),
      ),
    );
  }
}
