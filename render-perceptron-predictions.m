#!/usr/bin/env octave

hold all

setosa_label     = -1;
versicolor_label = +1;


% render training set
training_set = csvread('data/training-set');

petals = training_set(:, 1);
sepals = training_set(:, 2);

setosa_petals = petals(training_set(:, 3) == setosa_label, :);
setosa_sepals = sepals(training_set(:, 3) == setosa_label, :);

versicolor_petals = petals(training_set(:, 3) == versicolor_label, :);
versicolor_sepals = sepals(training_set(:, 3) == versicolor_label, :);

h = plot(
    setosa_petals,     setosa_sepals,     'or',
    versicolor_petals, versicolor_sepals, 'xb'
);
legend([h], {'setosa', 'versicolor'});


% render predictions
predictions = csvread('data/perceptron-predictions');

petals      = predictions(:, 1);
petal_count = length(unique(petals));

sepals      = predictions(:, 2);
sepal_count = length(unique(sepals));

shape = [sepal_count, petal_count];

petals = reshape(petals, shape);
sepals = reshape(sepals, shape);
labels = reshape(predictions(:, 3), shape);

colors = [
    hex2dec('FD'), hex2dec('B2'), hex2dec('B3')
    hex2dec('81'), hex2dec('B9'), hex2dec('DC')
];
colors = colors / 255;
colormap(colors);

contourf(petals, sepals, labels, 'LineColor', 'none');

hold off

axis ([
    min(petals) - 1,
    max(petals) + 1,
    min(sepals) - 1,
    max(sepals) + 1,
])

xlabel('Petal length, cm');
ylabel('Sepal length, cm');

title ('Perceptron predictions');
legend('location', 'southeast');

saveas(1, 'img/perceptron-predictions.png');
