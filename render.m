#!/usr/bin/env octave

training_set            = csvread('data/training-set');
training_set_setosa     = training_set(training_set(:, 3) == -1, [1:2]);
training_set_versicolor = training_set(training_set(:, 3) ==  1, [1:2]);

hold all

x  = training_set_setosa(:, 1);
y  = training_set_setosa(:, 2);
scatter(x, y, 'r', 'o');

x  = training_set_versicolor(:, 1);
y  = training_set_versicolor(:, 2);
scatter(x, y, 'b', 'x');

axis ([
    min(training_set(:, 1)) - 1,
    max(training_set(:, 1)) + 1,
    min(training_set(:, 2)) - 1,
    max(training_set(:, 2)) + 1,
])

xlabel('Petal length, cm');
ylabel('Sepal length, cm');
title ('Training set');
legend('location', 'southeast');
legend({'setosa', 'versicolor'});

hold off
saveas(1, 'img/training-set.png');

%

perceptron_predictions = csvread('data/perceptron-predictions');

petal  = perceptron_predictions(:, 1);
sepal  = perceptron_predictions(:, 2);
labels = perceptron_predictions(:, 3);

petal_points = length(unique(petal));
sepal_points = length(unique(perceptron_predictions(:, 2)));

petal  = reshape(petal,  [sepal_points, petal_points]);
sepal  = reshape(sepal,  [sepal_points, petal_points]);
labels = reshape(labels, [sepal_points, petal_points]);

colors = [
    hex2dec('FD'), hex2dec('B2'), hex2dec('B3')
    hex2dec('81'), hex2dec('B9'), hex2dec('DC')
];
colors = colors / 255;
colormap(colors);

clf
hold all
contourf(petal, sepal, labels, 'LineColor', 'none');

setosa_petals     = training_set_setosa(:, 1);
setosa_sepals     = training_set_setosa(:, 2);
versicolor_petals = training_set_versicolor(:, 1);
versicolor_sepals = training_set_versicolor(:, 2);

h = plot(
    setosa_petals,     setosa_sepals,     'or',
    versicolor_petals, versicolor_sepals, 'xb'
);

xlabel('Petal length, cm');
ylabel('Sepal length, cm');
title ('Perceptron predictions');
legend('location', 'southeast');
legend([h], {'setosa', 'versicolor'});

hold off
saveas(1, 'img/perceptron-predictions.png');
