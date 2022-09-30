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

import 'dart:math' as math;

import 'package:expansion_widget/expansion_widget.dart';
import 'package:flutter/material.dart';
import 'package:playground/constants/sizes.dart';
import 'package:playground/modules/examples/components/examples_components.dart';

const _kInitiallyExpanded = false;

class TypeFilter extends StatelessWidget {
  const TypeFilter({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(
        horizontal: kLgSpacing,
        vertical: kMdSpacing,
      ),
      child: ExpansionWidget(
        expandedAlignment: Alignment.topLeft,
        initiallyExpanded: _kInitiallyExpanded,
        titleBuilder: (animationValue, _, isExpanded, toggleFunction) =>
            InkWell(
          onTap: () => toggleFunction(animated: true),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              SizedBox(
                height: kContainerHeight,
                child: Row(
                  children: const <CategoryBubble>[
                    CategoryBubble(name: 'Example'),
                    CategoryBubble(name: 'Kata'),
                    CategoryBubble(name: 'Unit Test'),
                  ],
                ),
              ),
              Transform.rotate(
                angle: -math.pi * animationValue,
                alignment: Alignment.center,
                child: const Icon(Icons.keyboard_arrow_down),
              ),
            ],
          ),
        ),
        content: Padding(
          padding: const EdgeInsets.symmetric(vertical: kMdSpacing),
          child: Wrap(
            runSpacing: kMdSpacing,
            children: _buildItems(),
          ),
        ),
      ),
    );
  }

  List<Widget> _buildItems() {
    List<String> tags = ['Long', 'Short', 'Easy', 'Hard', 'Medium'];
    List<Widget> items = [];
    for (var tag in tags) {
      items.add(CategoryBubble(name: tag));
    }

    return items;
  }
}
