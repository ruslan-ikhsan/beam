/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import 'package:flutter/material.dart';
import 'package:playground_components/playground_components.dart';

class ExampleSelectorState with ChangeNotifier {
  final PlaygroundController _playgroundController;
  String _filterText;
  List<CategoryWithExamples> categories;
  List<String> selectedTags = [];

  ExampleSelectorState(
    this._playgroundController,
    this.categories, [
    this._filterText = '',
  ]);

  String get filterText => _filterText;

  addSelectedTag(String tag) {
    selectedTags.add(tag);
    notifyListeners();
  }

  removeSelectedTag(String tag) {
    selectedTags.remove(tag);
    notifyListeners();
  }

  void setFilterText(String text) {
    _filterText = text;
    notifyListeners();
  }

  void setCategories(List<CategoryWithExamples> categories) {
    this.categories = categories;
    notifyListeners();
  }

  void sortCategories() {
    final categories = _playgroundController.exampleCache.getCategories(
      _playgroundController.sdk,
    );

    final sortedCategories = categories
        .map((category) => CategoryWithExamples(
            title: category.title,
            examples: _sortCategoryExamples(category.examples)))
        .where((category) => category.examples.isNotEmpty)
        .toList();
    setCategories(sortedCategories);
  }

  List<ExampleBase> _sortCategoryExamples(List<ExampleBase> examples) {
    if (selectedTags.isEmpty && filterText.isEmpty) {
      return examples;
    }
    if (selectedTags.isNotEmpty && filterText.isEmpty) {
      return sortExamplesByTags(examples);
    }
    if (selectedTags.isEmpty && filterText.isNotEmpty) {
      return sortExamplesByName(examples);
    }
    final sorted = sortExamplesByTags(examples);
    return sortExamplesByName(sorted);
  }

  List<ExampleBase> sortExamplesByTags(List<ExampleBase> examples) {
    List<ExampleBase> sorted = [];
    for (var example in examples) {
      if (example.tags.toSet().containsAll(selectedTags)) {
        sorted.add(example);
      }
    }
    return sorted;
  }

  List<ExampleBase> sortExamplesByType(
    List<ExampleBase> examples,
    ExampleType type,
  ) {
    return examples.where((element) => element.type == type).toList();
  }

  List<ExampleBase> sortExamplesByName(List<ExampleBase> examples) {
    return examples
        .where((example) =>
            example.name.toLowerCase().contains(filterText.toLowerCase()))
        .toList();
  }
}
